package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto"
)

var _ sdk.Tx = (*TxStd)(nil)

var _ ITx = (*TxTransfer)(nil)

type TxStd struct {
	ITxs      []ITx       `json:"itx"`      //ITx接口，将被具体Tx结构实例化
	Signature []Signature `json:"sigature"` //签名数组
	ChainID   string      `json:"chainid"`  //ChainID: 执行ITx.exec方法的链ID
	MaxGas    uint64      `json:"maxgas"`   //Gas消耗的最大值
}

type ITx interface {
}
type Signature struct {
	Pubkey    crypto.PubKey `json:"pubkey"`    //可选
	Signature []byte        `json:"signature"` //签名内容
	Nonce     uint64        `json:"nonce"`     //nonce的值
}

type Address []byte

type TransItem struct {
	Address Address `json:"addr"` // 账户地址
	QOS     uint64  `json:"qos"`  // QOS
	QSCs    uint64  `json:"qscs"` // QSCs
}

type TransItems []TransItem

type TxTransfer struct {
	Senders   TransItems `json:"senders"`   // 发送集合
	Receivers TransItems `json:"receivers"` // 接收集合
}
