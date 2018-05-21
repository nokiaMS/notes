
let Timer = require("timer.js")
let assert = require('assert')

import { getBasicTxSet } from "./resource"
import { NodeSuite } from "./node_testsuite";

const nodeIpList = ["http://192.168.1.26:8101", "http://192.168.1.26:8102","http://192.168.1.26:8103"]

let nodeSuiteList=[]

function initNodeSuiteList(nRepeatPerTick, nTickSec, nDurationMin) {  
    console.log("do initNodeSuiteList() ... from process: ", process.pid)
    let nNodes = nodeIpList.length
    let oneNodeSuite = null
    for(let i=0; i < nNodes ; i++){
        oneNodeSuite = new NodeSuite(nodeIpList[i], nTickSec, nDurationMin*60)
        buildTxTaskList(oneNodeSuite, nRepeatPerTick)  // init test case for every suite:
        nodeSuiteList.push(oneNodeSuite)
    }
    return nodeSuiteList
}

/* 
  可以把原先的 tx = new ContractTx， tx.createTx()  , Tx --> Task 这三个步骤放在一个地方完成！
  原先的 即 setUp() 做了两件事: tx.createTx()  , Tx --> Task
 */
    // 基本思想： 把 Tx 对象的创建从 for 1..nTaskNumber 循环中拿出来， 只创建一次。 
    // 这应该是没问题的！
    // 结果就是： 一个 Tx 实例 + nTaskNumber个回调（但是在 for 1..nTaskNumber 循环 给 Tx 实例的 cb参数是不一样的），
    //          就能满足： 能够区分出、正确地拿到每次执行的结果的！
function buildTxTaskList(oneNodeSuite, nRepeatPerTick) { 
        var web3 = oneNodeSuite.getWeb3();
        const endPointUrl = oneNodeSuite.getEndPointUrl();        
        var tasks = []
        const basic_tx_set = getBasicTxSet(web3)
        for (let i=0; i < nRepeatPerTick; i++) {

            basic_tx_set.forEach((tx) => {
                const task = (cb) => {

                    //执行交易。 参数是一个callback，当交易完成(不管成功还是失败)时，这个callback会被调用！
                    tx.callTx(
                        function(e, result){
                            // console.log("On TxDoneCb, this pointer is: ", this) // 'this' value is undefined
                            cb(null, endPointUrl + "[task-" + i + "]: " + result)
                    });

                }
                tasks.push(task) 
            });
        }
        oneNodeSuite.setTxTaskList(tasks)
        return tasks
}

function runTestOfNode(index){
    assert.ok(index >= 1, "node index should be 1-based！")
    var myNode = nodeSuiteList[index-1]
    console.log(`call run of nodeSuiteList[${index-1}]`)
    myNode.run()
}

module.exports={init: initNodeSuiteList, run: runTestOfNode}

