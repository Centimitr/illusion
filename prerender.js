/*
    Reference: https://www.mxgw.info/t/phantomjs-prerender-for-seo.html
    Thanks!
*/

// config
var REFRESH_TIMEOURT = 500;
var MIN_REFRESH_TIMEOUT = 50;
var MAX_REFRESH_TIMEOUT = 1000;
var PROCESS_TIMEOUT = 5000;
var MIN_PROCESS_TIMEOUT = 500;
var MAX_PROCESS_TIMEOUT = 15000;

// initial
var page = require('webpage').create();
var system = require('system');

// params : url [processTimeout] [pageRefreshTimeoutAfterResourceLoad] [mobileOrDesktopViewport]
var url = system.args[1];
// var paramParseInt = function(paramValue,min,max,default){
//     value = parseInt(paramValue);
//     if (value>max)||(value<min){
//         value = default
//     }
//     return value
// };
// var isValidParam = function(value){
//     return value&&value!='0'&&value!='null'&&value!='NULL'&&value!='nil'
// }
// isValidParam(system.args[2]) && (PROCESS_TIMEOUT = paramParseInt(system.args[2],MIN_PROCESS_TIMEOUT,MAX_PROCESS_TIMEOUT,PROCESS_TIMEOUT));
// isValidParam(system.args[3]) && (REFRESH_TIMEOURT = paramParseInt(system.args[3],MIN_REFRESH_TIMEOUT,MAX_REFRESH_TIMEOUT,REFRESH_TIMEOURT));
// system.args[4] == 'desktop' && (page.viewportSize = {width: 1440,height: 1024});
// system.args[4] == 'mobile' && (page.viewportSize = {width: 960,height: 1536});

var refreshTimer;
var processTimer;
var resCnt = 0;
var failCnt = 0;

page.onResourceRequested = function(req){
    resCnt++;
    clearTimeout(refreshTimer);
};
page.onResourceReceived = function (res) {
	// event will trigger when resource isn't fully received in chunk mode
    if (res.stage == 'end'){
    	resCnt--;
    	if (resCnt === 0){
        	refreshTimer = setTimeout(capture, REFRESH_TIMEOURT);
   		}
	}
};
page.onResourceTimeout = function(req){
    resCnt--;
    failCnt++;
};
page.onResourceError = function(err){
    resCnt--;
    failCnt++;
};

// run
page.open(url, function (status) {
    if (status !== 'success') {
        phantom.exit(-1);
    } else {
        processTimer = setTimeout(function(){
			//output
    		console.log(page.content);
   			clearTimeout(processTimer);
    		phantom.exit(failCnt);
        }, PROCESS_TIMEOUT);
    }
});