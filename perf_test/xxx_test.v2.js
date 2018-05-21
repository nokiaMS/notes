import {contract_ok_data, money_ok_data, ContractTx, MoneyTransferTx} from "./resource"

let Timer = require("timer.js")

const Web3 = require('web3');


export class NodeSuite {

    constructor(endPointUrl) {
        this.testCases=[]
        // this.web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:10001"));
        try {
            this.web3 = new Web3(new Web3.providers.HttpProvider(endPointUrl));
            // console.log("this.web3.personal: ", this.web3.personal)

            console.log("account: ", this.web3.eth.accounts[0])
            // this.web3.personal.unlockAccount("0xe353a5d9907cd806eac46a6476f3c464e91d8209","123456",3000);

            var coinbase = this.web3.eth.coinbase;
            console.log(coinbase);

            var balance = this.web3.eth.getBalance(coinbase);
            console.log(balance.toString(10));

            this.myTimer = new Timer({
                tick    : 1, // 秒单位！
                ontick  : this.onTick,
                onstart : this.onTimerStart,
                onstop  : this.onTimerStop, 
                onend   : this.onTimerEnd
              });

        } catch (e) {
            console.log("create Web3 error: ", e);
        }
    }

    onTick(ms) {
        console.log(ms + ' milliseconds left');
        // doOperationPerMinute(nodeList[0, 100, null, null, 50]);
        this.testCases.forEach(testCase => {
            testCase.callTx();
        });
    }

    onTimerStart() {
        console.log('timer started')
    }

    onTimerStop() {
        console.log('timer onstop')
    }
    
    onTimerEnd() {
        console.log('timer ended normally')
    }

    getWeb3() {
        return this.web3;
    }

    addTestCase(testCase) {
        this.testCases.push(testCase)        
    }

    setTestCases(casesList, nRepeats){

        if (caseList.length != nRepeats.length) {
            console.error("参数1，参数2的长度必须相同！")
        } return 

        this.caseList = casesList
        this.nRepeats = nRepeats
    }

    runPararel() {

    }

    /**
     * 初始化工作， 比如部署智能合约
     */
    setUp(){
        // console.log("testCaseSet: ", this.testCases)
        this.testCases.forEach(testCase => {
            // console.log("testCase: ", testCase);
            testCase.createTx();
        });
    }

    /**
     * 执行交易
     */
    runTestCase(nMin) {
        this.myTimer.start(nMin*60)
    }

}

let testSuite = new NodeSuite("http://192.168.1.26:8101")
let thisNodeWeb3 = testSuite.getWeb3()

testSuite.addTestCase(new ContractTx(1, contract_ok_data[0], thisNodeWeb3.eth.accounts[0], thisNodeWeb3))
testSuite.addTestCase(new MoneyTransferTx(2, money_ok_data[0], thisNodeWeb3))
testSuite.addTestCase(new ContractTx(3, contract_ok_data[0], thisNodeWeb3.eth.accounts[0], thisNodeWeb3))

testSuite.setUp() 
testSuite.runTestCase(1)





