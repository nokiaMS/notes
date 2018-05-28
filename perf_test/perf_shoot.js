#!/usr/bin/node

const Web3 = require('web3')
const config = require('./config.js')

if(typeof(process.argv[2]) == 'undefined') {
	console.log("Usage: " + process.argv[1] + " <index>")
	return
} else {
	var index = parseInt(process.argv[2])
	console.log("The process index is: " + index)
}

var procInfo = config.procList[index]
console.log(procInfo.url)

//if object web3 has not been created, then we create it. if not, we should re-use the old one.
if(typeof web3 !== 'undefined') {
	web3 = new Web3(web3.currentProvider)
} else {
	web3 = new Web3(new Web3.providers.HttpProvider(procInfo.url))
}

if(web3.isConnected() !== true) {
	console.log("Can not connect to web3 server.")
	return
}

//Connected to web3 server.
var eth = web3.eth
var personal = web3.personal
var totalTransactionCount = 0

//unlock account.
personal.unlockAccount(procInfo.fromAccount, procInfo.password, 60*60*100)  //100 hours

//send transaction.
//can not use async function as nonce error.  /* guoxu core.ErrNonceTooLow */
function sendTransactionBatch() {
	eth.sendTransaction({from:procInfo.fromAccount,to:procInfo.toAccount,value:"1"}, 
    		function(error, result){
    			if(error) {
        			console.error(error)
    			} else {
				totalTransactionCount += 1
        			console.log("Total Transaction count: " + totalTransactionCount)
    			}
			sendTransactionBatch()
		})
}

function runSmartContractFunc() {
	var executeStr = "global.contractInstance." + config.testFunction + ".sendTransaction({from:'" + procInfo.fromAccount + "', gas: 100000})"
	console.log(executeStr)
	var executeStrWithLog = "console.log(" + executeStr + ".toString())"
	eval(executeStrWithLog)
}

var testType = procInfo.testType
if(testType == 1) {
	console.log("Running money transfer case...")
	setInterval(sendTransactionBatch,procInfo.interval)
}

if(testType == 2) {
	console.log("Running smart contract case...")
	const hwContract = web3.eth.contract(JSON.parse(abi))
	global.contractInstance = hwContract.at(procInfo.contractAddress)
	//setInterval(runSmartContractFunc,procInfo.interval)
	sendTransactionBatch()
}

