import {contract_ok_data, money_ok_data, ContractTx, MoneyTransferTx} from "./resource"

const Web3 = require('web3');
var async = require('async')

var endPointUrl = "http://127.0.0.1:50004"
var web3;
try {
    web3 = new Web3(new Web3.providers.HttpProvider(endPointUrl));
    // console.log("web3.personal: ", web3.personal)

    console.log("account: ", web3.eth.accounts[0])
    // web3.personal.unlockAccount("0xe353a5d9907cd806eac46a6476f3c464e91d8209","123456",3000);

    var coinbase = web3.eth.coinbase;
    console.log(coinbase);

    var balance = web3.eth.getBalance(coinbase);
    console.log(balance.toString(10));

} catch (e) {
    console.log("create Web3 error: ", e);
}


function buildTaskList(nTaskNumber) { 
    var tasks = []
    for (let i=0; i < nTaskNumber; i++) {
        const tx = new MoneyTransferTx(2, money_ok_data[0], web3)
        tx.createTx()
        const task = (cb) => {
            tx.callTx(
                function(e, result){
                    cb(null, "task-" + i + ": " + result)
            });
        }
        tasks.push(task)
    }

    async.parallel(tasks, function(err,result){  
        if (err) {  
            console.log(err);  
        }  
        console.log(" async.result: ", result);  
    })  

 }

 buildTaskList(10)



