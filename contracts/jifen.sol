pragma solidity ^0.4.22;

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {
  function mul(uint256 a, uint256 b) internal constant returns (uint256) {
    uint256 c = a * b;
    assert(a == 0 || c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal constant returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn‘t hold
    return c;
  }

  function sub(uint256 a, uint256 b) internal constant returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal constant returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}

//管理者
contract owned {
    address public owner;

    function owned() public {
        owner = msg.sender;
    }

	// onlyOwner 函数修改器是一个合约属性，可以被继承，还能被重写。它用于在函数执行前检查某种前置条件
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

	//实现所有权转移
    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}

//必须是合约创建者才能添加积分 销毁积分
contract JiFen is owned {

   using SafeMath for uint256;

   function JiFen() public {}

   //事件机制测试
   event Instructor(bytes32 txid,bytes32 memberOID);

   struct TxDetail{
      bytes32 txid;  //交易编号
      bytes32 memberOID;  //交易者编号
      string companyOID; //所属租户
      string siteOID; //所属场站
      string createTime; //创建时间
      string updateTime; //更新时间
      string integral; //积分数
      string types_deductibleAmount_consumeAmount_goods_goodsQty; //类型  抵扣金额 消费金额 消费商品 商品数量
   }

    //查看对应账号的积分余额。 任何人都可以查到任何地址的余额，正如所有数据在区块链上都是公开的。
	mapping(bytes32 => uint256) public balanceOf;
	//用户交易明细
	mapping(bytes32 => bytes32[]) public user_txs;
	//交易细节
	mapping(bytes32 => TxDetail) public tx_details;

	//分配积分流程函数1、增加积分 2、增加交易明细 3、维护用户交易明细
	function distributeJiFen_flow(uint256 value,bytes32 txid,bytes32 memberOID,string companyOID,string siteOID,string createTime,string updateTime,string integral,string types_deductibleAmount_consumeAmount_goods_goodsQty) onlyOwner public returns (bool, string){
	   balanceOf[memberOID] = balanceOf[memberOID].add(value);

	   tx_details[txid] = TxDetail(txid,memberOID,companyOID,siteOID,createTime,updateTime,integral,types_deductibleAmount_consumeAmount_goods_goodsQty);
       emit Instructor(txid, memberOID);

	   user_txs[memberOID].push(txid);
	   return (true,"success");
	}

	//分配积分流程函数1、减少积分 2、增加交易明细 3、维护用户交易明细
	function deductJiFen_flow(uint256 value,bytes32 txid,bytes32 memberOID,string companyOID,string siteOID,string createTime,string updateTime,string integral,string types_deductibleAmount_consumeAmount_goods_goodsQty) onlyOwner public returns (bool, string){
	   require(balanceOf[memberOID] >= value);                // Check if the targeted balance is enough
       balanceOf[memberOID] = balanceOf[memberOID].sub(value);                        // Subtract from the targeted balance

	   tx_details[txid] = TxDetail(txid,memberOID,companyOID,siteOID,createTime,updateTime,integral,types_deductibleAmount_consumeAmount_goods_goodsQty);
       emit Instructor(txid, memberOID);

	   user_txs[memberOID].push(txid);
	   return (true,"success");
	}

    function getUser_txsLength(bytes32 memberOID) public constant returns (uint256) {
        return user_txs[memberOID].length;
    }

}