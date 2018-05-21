import {ContractTx, MoneyTransferTx} from "./basic_tx_type"
export {ContractTx, MoneyTransferTx}

var seed = 1;
function random() {
    var x = Math.sin(seed++) * 1000000;
    return x - Math.floor(x);
}
export const money_ok_data=[
    {adrr:{hashAddrSrc:'', hashAddrDest:''}}
]

export const contract_ok_data=[
    {
        abi:[
            {
            "constant": false,
            "inputs": [{
                "name": "uAddr",
                "type": "address"
            }, {
                "name": "name",
                "type": "string"
            }],
            "name": "CreateUser",
            "outputs": [{
                "name": "",
                "type": "bool"
            }],
            "payable": false,
            "stateMutability": "nonpayable",
            "type": "function"
        }, {
            "constant": true,
            "inputs": [{
                "name": "",
                "type": "address"
            }],
            "name": "mapStoremanGroup",
            "outputs": [{
                "name": "",
                "type": "string"
            }],
            "payable": false,
            "stateMutability": "view",
            "type": "function"
        }, {
            "inputs": [],
            "payable": false,
            "stateMutability": "nonpayable",
            "type": "constructor"
        }, {
            "anonymous": false,
            "inputs": [{
                "indexed": false,
                "name": "add",
                "type": "address"
            }, {
                "indexed": false,
                "name": "name",
                "type": "string"
            }],
            "name": "EventCreateUser",
            "type": "event"
        }],
        bin:'0x6060604052341561000f57600080fd5b6104478061001e6000396000f30060606040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680637af318d514610051578063e651a4a9146100e5575b600080fd5b341561005c57600080fd5b6100cb600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610197565b604051808215151515815260200191505060405180910390f35b34156100f057600080fd5b61011c600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506102c6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561015c578082015181840152602081019050610141565b50505050905090810190601f1680156101895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906101eb929190610376565b507fa7507b65bddb7fb6efe748528d63193dacd81d9a2e8d134dbc80550fbb23a4678383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610281578082015181840152602081019050610266565b50505050905090810190601f1680156102ae5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16001905092915050565b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561036e5780601f106103435761010080835404028352916020019161036e565b820191906000526020600020905b81548152906001019060200180831161035157829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103b757805160ff19168380011785556103e5565b828001600101855582156103e5579182015b828111156103e45782518255916020019190600101906103c9565b5b5090506103f291906103f6565b5090565b61041891905b808211156104145760008160009055506001016103fc565b5090565b905600a165627a7a723058206c80d6e166ed54ab46a4a144908cbb5cea2548d83f394bfb48fc72e4852a14b70029',
        gas:'4700000'
    }
]

export const contract_err_data=[]

export function getBasicTxSet_Ok(_web3) {
    var basic_tx_set=[]    
    const tx = new MoneyTransferTx(2, money_ok_data[0], _web3)
    tx.createTx()
    basic_tx_set.push(tx)
    return basic_tx_set
}

export function getBasicTxSet(web3) {
    var basic_tx_set=[]    
    const accounts = web3.eth.accounts

    let inData = {addr:{hashAddrSrc: accounts[0], hashAddrDest:accounts[1]}}
    let tx = new MoneyTransferTx(1, inData, web3)
    tx.createTx()
    basic_tx_set.push(tx)

    inData = {addr:{hashAddrSrc: accounts[1], hashAddrDest:accounts[0]}}
    tx = new MoneyTransferTx(2, inData, web3)
    tx.createTx()
    basic_tx_set.push(tx)

    tx = new ContractTx(3, contract_ok_data[0], accounts[0], web3)
    basic_tx_set.push(tx)

    return basic_tx_set
}


console.log(MoneyTransferTx)


