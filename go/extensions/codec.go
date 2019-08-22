package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterInterface((*ITx)(nil), nil)
	codec.RegisterConcrete(TxCreateMarket{}, "microtick/CreateMarket", nil)
	codec.RegisterConcrete(&Signature{}, "qbase/txs/signature", nil)
	codec.RegisterConcrete(&TxStd{}, "qbase/txs/stdtx", nil)
	codec.RegisterConcrete(&TxTransfer{}, "qos/txs/TxTransfer", nil)
}
