package rest

import (
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/common"

)
var cdc = amino.NewCodec()

type Signature interface {
	Bytes() []byte
	IsZero() bool
	Equals(Signature) bool
}

func SignatureFromBytes(pubKeyBytes []byte) (pubKey Signature, err error) {
	err = cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}

type MsgRegisterVpnService struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Ip            string `json:"ip"`
	UploadSpeed   int64  `json:"upload_speed"`
	DownloadSpeed int64  `json:"download_speed"`
	Ppgb          int64  `json:"price_per_gb"`
	EncMethod     string `json:"enc_method"`
	Latitude      int64  `json:"location_latitude"`
	Longitude     int64  `json:"location_longitude"`
	City          string `json:"location_city"`
	Country       string `json:"location_country"`
	NodeType      string `json:"node_type"`
	Version       string `json:"version"`
}
type MsgRegisterMasterNode struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
}

type MsgDeleteVpnUser struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Address  string `json:"address", omitempty`
}
type MsgDeleteMasterNode struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Address  string `json:"address"`
}
type MsgPayVpnService struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Coins        string `json:"amount", omitempty`
	Vpnaddr      string `json:"vaddress", omitempty`
	SigName      string `json:"sig_name"`
	SigPassword  string `json:"sig_password"`
}
type MsgGetVpnPayment struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Coins        string `json:"amount"`
	Sessionid    string `json:"session_id"`
	Counter      int64  `json:"counter"`
	IsFinal      bool   `json:"isfinal"`
	Signature    string `json:"sign"`
}

type MsgRefund struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Sessionid string `json:"session_id", omitempty`
}

type ClientSignature struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	Coins        string `json:"amount"`
	Sessionid    string `json:"session_id"`
	Counter      int64  `json:"counter"`
	IsFinal      bool   `json:"isfinal"`
}

type Response struct {
	Success bool            `json:"success"`
	Hash    string          `json:"hash"`
	Height  int64           `json:"height"`
	Data    []byte          `json:"data"`
	Tags    []common.KVPair `json:"tags"`
}

type SendTokens struct {
	BaseReq 	  utils.BaseReq `json:"base_req"`
	ToAddress string `json:"to"`
	Coins     string `json:"amount"`
}