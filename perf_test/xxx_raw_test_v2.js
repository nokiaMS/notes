import {contract_ok_data, money_ok_data, ContractTx, MoneyTransferTx, getBasicTxSet} from "./resource"

const Web3 = require('web3');
var async = require('async')

var endPointUrl = "http://192.168.1.26:8101"

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

// 基本思想： 把 Tx 对象的创建从 for 1..nTaskNumber 循环中拿出来， 只创建一次。 
// 这应该是没问题的！
// 结果就是： 一个 Tx 实例 + nTaskNumber个回调（但是在 for 1..nTaskNumber 循环 给 Tx 实例的 cb参数是不一样的），
//          就能满足： 能够区分出、正确地拿到每次执行的结果的！

function buildTaskList(nTaskNumber) { 
    var tasks = []
    const basic_tx_set = getBasicTxSet(web3)
 
    for (let i=0; i < nTaskNumber; i++) {

        basic_tx_set.forEach((tx) => {
            const task = (cb) => {
                tx.callTx(
                    function(e, result){
                        cb(null, "task-" + i + ": " + result)
                });
            }
            tasks.push(task) 
        });
    }

    async.parallel(tasks, function(err,result){  
        if (err) {  
            console.log(err);  
        }  
        console.log(" async.result: ", result);  
    })  

 }

 buildTaskList(10)



