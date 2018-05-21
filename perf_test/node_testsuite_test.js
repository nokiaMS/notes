import {contract_ok_data, money_ok_data, ContractTx, MoneyTransferTx} from "./resource"
import {NodeSuite} from "./node_testsuite"

/**
 * 下面的代码要移除掉！！！ 否则在被 required 的时候就会执行！
 */
let testSuite = new NodeSuite("http://192.168.1.26:8101", 2, 10)   // 持续运行的时间长度，单位：秒
let thisNodeWeb3 = testSuite.getWeb3()

console.log("testSuite.web3.eth.accounts: ", thisNodeWeb3.eth.accounts)

testSuite.addTestCase(new ContractTx(1, contract_ok_data[0], thisNodeWeb3.eth.accounts[0], thisNodeWeb3))
testSuite.addTestCase(new MoneyTransferTx(2, money_ok_data[0], thisNodeWeb3))
testSuite.addTestCase(new ContractTx(3, contract_ok_data[0], thisNodeWeb3.eth.accounts[0], thisNodeWeb3))

testSuite.setUp()

testSuite.run()
