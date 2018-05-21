
export class ContractTx {

    constructor(id, contractInfo, from, web3) {
        this.id = id
        this.TxObj = null
        this.contract=null
        this.inputData = contractInfo
        this.from = from
        this.web3 = web3
        let nowDate = new Date();
        this.name_prefix = "name123"+ nowDate.toLocaleDateString() + " "+ nowDate.toLocaleTimeString();
    }

    // 创建智能合约
    createTx(cb) {

        this.TxObj = this.web3.eth.contract(this.inputData.abi).new({
            from: this.from,
            data: this.inputData.bin,
            gas: '4700000'
        }, function (e, contract) {
            console.log("ContractTx::createTx() result: ", e, contract);
            if (typeof contract.address !== 'undefined') {
                console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
                this.contract = contract;
            }
            if (cb !== undefined) {
                cb(e, contract);
            }
        })

    }

    // 调用智能合约
    callTx(cb) {
        if (this.TxObj && this.contract) {
            this.TxObj.CreateUser(this.from, this.name_prefix + random())   
        }
        cb(null, 'ContractTx') // 一律返回成功
    }

}


export class MoneyTransferTx {

    /**
     * 
     * @param {*} id 
     * @param {*} transferInfo : 为  {adrr:{hashAddrSrc:'', hashAddrDest:''}} 结构。如果没有给出hashAddrSrc、hashAddrDest，那么默认是节点内的帐号之间转账
     * @param {*} web3 
     */
    constructor(id, transferInfo, web3) {
        this.id = id;
        this.inputData = transferInfo;
        this.web3 = web3; 
        console.assert(this.web3 !== undefined, "web3参数不能为undefined")
    }

    // 转账交易
    createTx(cb) {
        ;
    }

    // 执行转账：
    callTx(cb) {
        let src=""
        let dest=""
        let accounts=this.web3.eth.accounts;

        console.assert(this instanceof MoneyTransferTx, "this指针不正确！")

        src = (this.inputData && this.inputData.hashAddrSrc) ? this.inputData.hashAddrSrc : accounts[0];
        dest = (this.inputData && this.inputData.hashAddrDest) ? this.inputData.hashAddrDest : accounts[1];

        // console.log("srcAddr: ", src, ",  destAddr: ", dest)
        // console.log("type of srcAddr: ", typeof src, ",  typeof destAddr: ", typeof dest)
        this.web3.eth.sendTransaction({
                from:src,
                to:dest,
                value:this.web3.toWei(1,"ether")},function (error, result) {
                    //WYH: 这里为了不中断剩余的task，所以就一律用 null 作为 cb 的第一个参数： 
                    cb(null, "MoneyTxHash:" + result)
                })
    }

}
