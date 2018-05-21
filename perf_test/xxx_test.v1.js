import {contract_ok_data, ContractTx} from "./resource"

const Web3 = require('web3');

class TestCase_ContractTx {

    /**
     * 
     * @param {*} id 
     * @param {*} txInput : 比如合约的 {abi:, bin:, gas: } 对象
     * @param {*} web3 
     */
    constructor(id, txInput, web3){
        this.id=id
        this.contractTx = new ContractTx(txInput, web3.eth.accounts[0])

    }

    run(){
        console.log("testcase::run()... of ID: ", this.id)
        this.contractTx.callTx()
    }
}

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

        } catch (e) {
            console.log("create Web3 error: ", e);
        }
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

    runTestCase() {
        this.testCases.forEach(testCase => {
            testCase.run();
        });
    }

    run(nodeAddr, repeats, txMoney, txContract, nRatio){
   
        console.log("testNode is: ", TargetOps)
    
        let nTxMoney = repeats * nRatio / 100;
        let nTxContract = repeats - nTxMoney;
        if ( txMoney !== 'undefined') {
        for (let i=0; i < nTxMoney; i++) {
            // testNode.sendMoney()
            txMoney.callTx()
        }
        }
    
        if  (txContract !== 'undefined')
        for (let i=0; i < nTxContract; i++) {
            // testNode.sendContract()
            txMoney
        }    
    }
 
}

let testSuite = new NodeSuite("http://192.168.1.26:8101")

let oo = new TestCase_ContractTx(1, contract_ok_data[0], testSuite.getWeb3());



testSuite.addTestCase(new TestCase_ContractTx(1, contract_ok_data[0], testSuite.getWeb3()))
testSuite.addTestCase(new TestCase_ContractTx(2, contract_ok_data[0], testSuite.getWeb3()))
testSuite.addTestCase(new TestCase_ContractTx(3, contract_ok_data[0], testSuite.getWeb3()))

testSuite.runTestCase()





