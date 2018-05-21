var cluster = require('cluster');
var http = require('http');
var numCPUs = require('os').cpus().length;

var args = process.argv.splice(2);

var nRepeatPerTick
var nTickSec
var nDurationMin

if (cluster.isMaster) {
    if (args[0] === undefined || args[1] == undefined || args[2] == undefined) {
        console.log("Usage: node  app.js  nRepeatsPerTick  nTickSec  nDurationMin ")
        process.exit(0)
    }
    nRepeatPerTick = parseInt(args[0])
    nTickSec = parseInt(args[1])
    nDurationMin = parseInt(args[2])
}

if (cluster.isMaster) {
    console.log("detect out cpu numbers: ", numCPUs);
    console.log("master start...");
    console.log("nRepeat: ", nRepeatPerTick, " duration: ", nDurationMin)

    // Fork workers.
    numCPUs = 2;

    let my_env_vars = {test_nRepeats: nRepeatPerTick, test_nTickSec: nTickSec, test_nDurationMin: nDurationMin}
    for (var i = 0; i < numCPUs; i++) {
        cluster.fork(my_env_vars);
    }

    cluster.on('exit', function(worker, code, signal) {
        console.log('worker ' + worker.process.pid + ' died');
    });
} else {
    console.log(`worker env var: [test_nRepeats: ${process.env.test_nRepeats}, test_nDurationMin: ${process.env.test_nDurationMin}] `)
    
    //每一个worker都会加载一次 "./scheduler"
    let tester = require("./test_builder");
    
    let nodeSuiteList = tester.init(process.env.test_nRepeats, process.env.test_nTickSec, process.env.test_nDurationMin)

    console.log(`I am worker #${cluster.worker.id}`)

    let func_exit = function(nSeccond) {
        console.log(`process #{process.pid} will quit after ${nSeccond}s`); 
        process.send("exit");
        setTimeout(function() { 
            console.log('quitting'); 
            process.exit(0); 
        }, 
        nSeccond*1000);
    }

    var bTesting = false
    if (!bTesting) {
        tester.run(cluster.worker.id);
    } else {
        var Timer = require("timer.js")
        function mimic_jobs(){
            var myTimer = new Timer({
                tick    : 5,
                ontick  : function(ms) { console.log(ms + ' milliseconds left') },
                onstart : function() { console.log('timer started') },
                onstop  : function() { console.log('timer stop') },
                onpause : function() { console.log('timer set on pause') },
                onend   : function() { console.log('timer ended normally'); func_exit(1); }
              });
              myTimer.start(15)
         }
         mimic_jobs()
        //  func_exit(1) // 这个会导致process过早退出（timer还没执行完呢！ ）！
        
    }
}
