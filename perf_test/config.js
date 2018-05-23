//test process count

abi = '[{"constant":true,"inputs":[],"name":"sayHi","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]'
testFunction = 'sayHi'

var proc1 = {
	fromAccount: "0xe353a5d9907cd806eac46a6476f3c464e91d8209",
	password: "123456",
	toAccount: "0xeABB235cb2a269DBD1dFcd71d51A2F9EA3DC184B",
	url: "http://192.168.138.139:50001",
	contractAddress: '0x3b5e407420efbd6558fc1d6466822432d0fcf9cc',
	interval: 10,
	testType: 2 //1.transfer money; 2.contract call.
}

var proc2 = {
	fromAccount: "0xeAD0f7f9a5bC0E19AEbfBc0Fb1f12EB05e0664Ed",
	password:"123456",
	toAccount: "0xeABB235cb2a269DBD1dFcd71d51A2F9EA3DC184B",
	url: "http://192.168.138.139:50002",
	contractAddress: '0xdbf1b2c85d417239b0012ed6af1f866191b0d754',
	interval: 10,
	testType: 2 //1.transfer money; 2.contract call.
}

var proc3 = {
	fromAccount: "0xF03DC9b4d040bD667acB21d6ec1384343Bd1F45d",
	password:"123456",
	toAccount: "0xeABB235cb2a269DBD1dFcd71d51A2F9EA3DC184B",
	url: "http://192.168.138.139:50003",
	contractAddress: '0xec374ebb7dec3bfe878bf71636d139dfe284ac73',
	interval: 10,
	testType: 2 //1.transfer money; 2.contract call.
}

var proc4 = {
	fromAccount: "0x2D026DE651A42243cD01b25ecE54CD4511082817",
	password:"123456",
	toAccount: "0xeABB235cb2a269DBD1dFcd71d51A2F9EA3DC184B",
	url: "http://192.168.138.139:50004",
	contractAddress: '0x1a092195de3ce6073779ecb2d438364c90c3b653',
	interval: 10,
	testType: 2 //1.transfer money; 2.contract call.
}

var procList = []
procList.push(proc1)
procList.push(proc2)
procList.push(proc3)
procList.push(proc4)

module.exports.procList = procList
module.exports.testFunction = testFunction
