const fs = require('fs');
const solc = require('solc');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("http://192.168.138.139:50001"));
const input = fs.readFileSync('contract.sol');
const output = solc.compile(input.toString(), 1);
const bytecode = output.contracts[':HelloWorldContract'].bytecode;
const abi = output.contracts[':HelloWorldContract'].interface;
const helloWorldContract = web3.eth.contract(JSON.parse(abi));
console.log(abi)
console.log("Unlocking coinbase account.");
console.log("web3 version is: " + web3.version.api);
try {
	web3.personal.unlockAccount("0xe353a5d9907cd806eac46a6476f3c464e91d8209","123456",3000);
} catch(e) {
	console.log(e);
	return;
}

console.log("Deploying the contract.");

const helloWorldContractInstance = helloWorldContract.new({
	data: '0x' + bytecode,
	from: "0xe353a5d9907cd806eac46a6476f3c464e91d8209",
	gas: 200000
	}, (err, res) => {
	if(err) {
		console.log(err);
		return;
	}
	console.log(res.transactionHash);
	if(res.address) {
		console.log('Contract address: ' + res.address);
	}
});

