
let Timer = require("timer.js")
let async = require("async")
const Web3 = require('web3');


export class NodeSuite {

    /**
     * 
     * @param {*} endPointUrl : 测试Node的RPC地址
     * @param {*} tickSec : 每隔多长时间调度一次(秒单位)
     * @param {*} durationSec : 持续运行的时间长度(秒单位)
     */
    constructor(endPointUrl,  tickSec, durationSec) {
        this.tickSec    = tickSec
        this.durationSec = durationSec
        this.endPointUrl = endPointUrl
        // this.web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:10001"));
        try {
            this.web3 = new Web3(new Web3.providers.HttpProvider(endPointUrl));
            console.log("account: ", this.web3.eth.accounts[0])
            // this.web3.personal.unlockAccount("0xe353a5d9907cd806eac46a6476f3c464e91d8209","123456",3000);

            var coinbase = this.web3.eth.coinbase;
            console.log(coinbase);

            var balance = this.web3.eth.getBalance(coinbase);
            console.log(balance.toString(10));

            this.myTimer = new Timer({
                tick    : this.tickSec, // 秒单位！
                ontick  : (ms) => { this.onTick(ms) },
                onstart : () => { this.onTimerStart() },
                onstop  : () => { this.onTimerStop() }, 
                onend   : () => { this.onTimerEnd() }
              });

        } catch (e) {
            console.log("create Web3 error: ", e);
        }
    }

    getEndPointUrl() {return this.endPointUrl;}
    getWeb3() { return this.web3;}

    //=========================  Timer 相关函数 =========================
    onTick(ms) {
        // 使用箭头函数后， 这里的this 就指向 NodeSuite 实例了！
        console.log(ms + ' milliseconds left @', (new Date()).toLocaleTimeString());
        this.runTxTaskList()
    }

    onTimerStart() {
        console.log('timer started @', (new Date()).toLocaleTimeString())
        this._startBlockNumber = this.web3.eth.blockNumber
        this.runTxTaskList()
    }

    onTimerStop() {
        console.log('timer onstop @', (new Date()).toLocaleTimeString())
    }
    
    onTimerEnd() {
        console.log('timer ended normally @', (new Date()).toLocaleTimeString())
        let endBlockNumber = this.web3.eth.blockNumber
        let averageMiningSpeed = (endBlockNumber - this._startBlockNumber) / (this.durationSec / 60)
        console.log("平均出块速度: ", averageMiningSpeed)
        process.exit(0)
    }
    
    //=========================  Timer 相关函数 end =========================
    setTxTaskList(tasksList) {
        this.tasksList = tasksList
    }

    runTxTaskList() {
        console.log("runTxTaskList... tasksList length: ", this.tasksList.length)
        async.parallel(this.tasksList,function(err,result){  
            if (err) {
                console.log(err);  
            } 
            console.log("runTxTaskList result: ", result);  
            console.log("runTxTaskList end");  
        })  
    }

    /**
     * 初始化工作， 比如部署智能合约
     */
    setUp(){ 
    }

    /**
     * 开始测试： 
     */
    run() {
        console.log("run Node of url: ", this.endPointUrl)
        if (this.myTimer) {
            console.log(`启动定时器.... 定时时间:  ${this.durationSec} 秒， 周期: ${this.tickSec} 秒 `)
            this.myTimer.start(this.durationSec)
        }else{
            console.log("timer初始化失败，请检查原因")
        }
    }
}




