#!/usr/bin/node

const Web3 = require('web3')

if(typeof(process.argv[2]) == 'undefined') {
	console.log("Usage: " + process.argv[1] + " <url>")
	return
}

var url = process.argv[2]
console.log("Connecting to " + url)

//if object web3 has not been created, then we create it. if not, we should re-use the old one.
if(typeof web3 !== 'undefined') {
	web3 = new Web3(web3.currentProvider)
} else {
	web3 = new Web3(new Web3.providers.HttpProvider(url))
}

if(web3.isConnected() !== true) {
	console.log("Can not connect to web3 server.")
	return
}

//Connected to web3 server.
var eth = web3.eth
var oldBlockNum = eth.blockNumber

//Calculate the transaction count of per minutes.
function calcTranCount(startBlockNum, endBlockNum) {
	var totalTranCount = 0.0000
	var tranPerMin = 0.000000
	for(var i = startBlockNum + 1; i <= endBlockNum; i++) {
		totalTranCount += eth.getBlockTransactionCount(i)
		//console.log("Total transaction count: " + totalTranCount/min)
	}
	//get total time
	startTime = eth.getBlock(startBlockNum)['timestamp']
	endTime = eth.getBlock(endBlockNum)['timestamp']
	tranPerSec = (totalTranCount * 60 )/ (endTime - startTime)
	console.log("    Blocks:[from: " + startBlockNum + " to: " + endBlockNum + "] tranPerMin: " + tranPerMin+ " sec: " + (endTime - startTime))
}

//Show the result of one minutes.
function getResult() {
	var curBlockNum = eth.blockNumber
	console.log("Blocks per minute: " + (curBlockNum - oldBlockNum))
	calcTranCount(oldBlockNum, curBlockNum)
	oldBlockNum = curBlockNum	
}

//Get the count of blocked per minute.
setInterval(getResult,60*1000)

