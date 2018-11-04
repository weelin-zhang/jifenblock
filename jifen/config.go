package jifen

type Account struct {
	Address string
}

var JifenConfig *JifenConf

type JifenConf struct {
	LogPath      string
	LogLevel     string
	KeyDir       string
	HostUrl      string
	OwnerAccount string
	OwerSecret   string
	ContractAddr string
	AccountList  []string
	JifenSession *JifenSession
	NeedUpdateAccount chan string
}



type TxDetail struct {
	Id                                             int
	TxId                                           string `orm:"column(txid)"`      // 交易编号(32字符)
	TxType                                         int    `orm:"column(txtype)"`    //交易类型1分配积分,0扣除积分
	MemberOID                                      string `orm:"column(memberid)"`  // 用户账号
	CompanyOID                                     string `orm:"column(companyid)"` // 所属租户
	SiteOID                                        string `orm:"column(siteid)"`    // 所属场站
	CreateTime                                     string `orm:"column(created)"`   // 创建时间
	UpdateTime                                     string `orm:"column(updated)"`   // 更新时间
	Integral                                       int64  `orm:"column(integral)"`  // 积分数
	TpesDeductibleAmountConsumeAmountGoodsGoodsQty string `orm:"column(tdagq)"`     // 类型  抵扣金额 消费金额 消费商品 商品数量
}

