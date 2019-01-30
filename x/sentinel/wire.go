package sentinel

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterWire(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterVpnService{}, "sentinel/registervpn", nil)
	//cdc.RegisterConcrete(MsgDeleteVpnUser{}, "sentienl/deletevpnservice", nil)
	//cdc.RegisterConcrete(MsgRegisterMasterNode{}, "sentinel/masternoderegistration", nil)
	//cdc.RegisterConcrete(MsgDeleteMasterNode{}, "sentinel/deletemasternode", nil)
	//cdc.RegisterConcrete(MsgPayVpnService{}, "sentinel/payvpnservice", nil)
	//cdc.RegisterConcrete(MsgRefund{}, "sentinel/clientrefund", nil)
	//cdc.RegisterConcrete(MsgGetVpnPayment{}, "sentinel/getvpnpayment", nil)
	//cdc.RegisterConcrete(MsgSendTokens{}, "sentinel/sendtoken", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterWire(msgCdc)
	codec.RegisterCrypto(msgCdc)
}
