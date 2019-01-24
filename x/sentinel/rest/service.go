package rest

import (
	"encoding/json"
	"fmt"

	"net/http"
	"reflect"
	"strconv"
	"strings"

	ioutill "io/ioutil"

	"github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/sentinel"
)

/**
* @api {get} /keys/seed To get seeds for generate keys.
* @apiName getSeeds
* @apiGroup Sentinel-Tendermint
* @apiSuccessExample Response:
{
* garden sunset night final child popular fall ostrich amused diamond lift stool useful brisk very half rice evil any behave merge shift ring chronic
* }
*/
/**
* @api {post} /keys To get account.
* @apiName getAccount
* @apiGroup Sentinel-Tendermint
* @apiParam {String} name Name Account holder name.
* @apiParam {String} password Password password for account.
* @apiParam {String} seed Seed seed words to get account.
* @apiError AccountAlreadyExists AccountName is  already exists
* @apiError AccountSeedsNotEnough Seed words are not enough
* @apiErrorExample AccountAlreadyExists-Response:
* {
*   Account with name XXXXX... already exists.
* }
* @apiErrorExample AccountSeedsNotEnough-Response:
* {
*  recovering only works with XXX word (fundraiser) or 24 word mnemonics, got: XX words
* }
* @apiSuccessExample Response:
*{
*    "name": "vpn",
*    "type": "local",
*    "address": "cosmosaccaddr1udntgzszesn7z3xm64hafvjlegrh38ukzw9m7g",
*    "pub_key": "cosmosaccpub1addwnpepqfjqadxwa9p8tvwhydsakyvkajxgyd0ycanv25u7yff7lqtkwuk8vqcy5cg",
*    "seed": "hour cram bike donor script fragile together derive capital joy glance morning athlete special hint scrub guitar view popular dream idle inquiry transfer often"
*}
 */
/**
* @api {post} /register/vpn To register VPN service provider.
* @apiName registerVPN
* @apiGroup Sentinel-Tendermint
* @apiParam {String} ip Ip address of VPN service provider.
* @apiParam {Number} upload_speed Upload Net speed of VPN service.
* @apiParam {Number} download_speed Download Net speed of VPN service.
* @apiParam {Number} price_per_gb Price per GB.
* @apiParam {String} enc_method Encryption method.
* @apiParam {Number} location_latitude  Latitude Location of service provider.
* @apiParam {Number} location_longitude  Longiude Location of service provider.
* @apiParam {String} location_city  City Location of service provider.
* @apiParam {String} location_country  Country Location of service provider.
* @apiParam {String} node_type  Node type.
* @apiParam {String} version version.
* @apiParam {String} name Account name of service provider.
* @apiParam {string} password password of account.
* @apiParam {Number} gas Gas value.
* @apiError AccountAlreadyExists VPN service provider already exists
* @apiError NetSpeedInvalidError Netspeed is Invalid
* @apiError IpAddressInvalidError IP address is Invalid
* @apiError Price_per_GBInvalidError Price per GB is Invalid
* @apiErrorExample AccountAlreadyExists-Response:
*{
 * checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      13
* ABCICode:  1245197
* Error:     --= Error =--
* Data: common.FmtError{format:"Address already Registered as VPN node", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiErrorExample NetSpeedInvalidError-Response:
*{
 * checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      13
* ABCICode:  1245197
* Error:     --= Error =--
* Data: common.FmtError{format:"NetSpeed is not Valid", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiErrorExample IpAddressInvalidError-Response:
*{
 * "  invalid Ip address."
*}
* @apiErrorExample Price_per_GBInvalidError-Response:
*{
 * checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      13
* ABCICode:  1245197
* Error:     --= Error =--
* Data: common.FmtError{format:"Price per GB is not Valid", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiSuccessExample Response:
*{
*   "Success": true,
*   "Hash": "CF8E073D624F7FA6A41C3CAD9B4A1DB693234225",
*   "Height": 343,
*   "Data": "eyJ0eXBlIjoic2VudGluZWwvcmVnaXN0ZXJ2cG4iLCJ2YWx1ZSI6eyJGc3BlZWQiOiIxMiIsIlBwZ2IiOiIyMyIsIkxvY2F0aW9uIjoiaHlkIn19",
*    "Tags": [
*        {
*            "key": "dnBuIHJlZ2lzdGVyZWQgYWRkcmVzcw==",
*            "value": "Y29zbW9zYWNjYWRkcjFlZ3RydjdxdGU0NnY2cXEzN3p0YzB2dzRuMmhrejZuempycDVhZQ=="
*        }
*		    ]
*}
*/

func NewResponse(success bool, hash string, height int64, data []byte, tags []common.KVPair) Response {
	//var res Response
	return Response{
		Success: success,
		Height:  height,
		Hash:    hash,
		Data:    data,
		Tags:    tags,
	}
}
func registervpnHandlerFn(ctx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var a int64
		msg := MsgRegisterVpnService{}
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		}

		if !validateIp(msg.Ip) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid Ip address."))
			return

		}
		if reflect.TypeOf(msg.Ppgb) != reflect.TypeOf(a) || msg.Ppgb < 0 {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid amount of price per Gb"))
			return
		}
		if msg.UploadSpeed <= 0 || msg.DownloadSpeed <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid net speed details"))
			return
		}
		if msg.Latitude <= -90*10000 || msg.Longitude <= -180*10000 || msg.Latitude > 90*10000 || msg.Longitude > 180*10000 || msg.City == "" || msg.Country == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid  Location details"))
			return
		}
		if msg.NodeType == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Node type is required"))
			return
		}
		if msg.Version == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Version is required"))
			return
		}

		//ctx = ctx.WithGas(msg.Gas)
		//ctx = ctx.WithFromAddressName(msg.Localaccount)


		//fmt.Printf("Name content: ",ctx.GetFromName())
		//fmt.Printf("Address content: ",ctx.GetFromAddress())
		addr, err := ctx.GetFromAddress()
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}

		acc, err := ctx.GetAccountNumber(addr)
		seq, err := ctx.GetAccountSequence(addr)
		fmt.Printf("seq content: ",seq)
		//ctx = ctx.WithSequence(seq)
		//ctx = ctx.WithAccountNumber(acc)
		if err != nil {
			w.Write([]byte("account number error"))
			w.Write([]byte(string(acc)))

		}

		msg1 := sentinel.NewMsgRegisterVpnService(addr, msg.Ip, msg.Ppgb, msg.UploadSpeed, msg.DownloadSpeed, msg.EncMethod, msg.Latitude, msg.Longitude, msg.City, msg.Country, msg.NodeType, msg.Version)
		fmt.Printf("msg1 content: ",msg1)
		//txBytes, err := ctx.SignAndBuild(msg.Localaccount, msg.Password, []sdk.Msg{msg1}, cdc)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		//res, err := ctx.BroadcastTx(txBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		//respon := NewResponse(true, res.Hash.String(), res.Height, res.DeliverTx.Data, res.DeliverTx.Tags)
		//data, err := json.MarshalIndent(respon, "", " ")
		//w.Write(data)
	}
	return nil
}


/**
* @api {post} /register/master To register Master Node.
* @apiName registerMasterNode
* @apiGroup Sentinel-Tendermint
* @apiParam {String} name  Account name of Master Node.
* @apiParam {Number} gas Gas value.
* @apiParam {string} password password of account.
* @apiError AccountAlreadyExists Master Node already exists
* @apiErrorExample AccountAlreadyExists-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      13
* ABCICode:  1245197
* Error:     --= Error =--
* Data: common.FmtError{format:"Address already Registered as VPN node", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiSuccessExample Response:
{
*{
 *   "Success": true,
*    "Hash": "CF8E073D624F7FA6A41C3CAD9B4A1DB693234225",
*    "Height": 343,
*    "Data": "eyJ0eXBlIjoic2VudGluZWwvcmVnaXN0ZXJ2cG4iLCJ2YWx1ZSI6eyJGc3BlZWQiOiIxMiIsIlBwZ2IiOiIyMyIsIkxvY2F0aW9uIjoiaHlkIn19==",
*    "Tags": [
*        {
*             "key": "dnBuIHJlZ2lzdGVyZWQgYWRkcmVzcw==",
*             "value": "Y29zbW9zYWNjYWRkcjFlZ3RydjdxdGU0NnY2cXEzN3p0YzB2dzRuMmhrejZuempycDVhZQ=="
*         }
*             ]
* }
*/
//func registermasterdHandlerFn(ctx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		msg := MsgRegisterMasterNode{}
//		var err error
//		body, err := ioutill.ReadAll(r.Body)
//		if err != nil {
//			return
//		}
//
//		json.Unmarshal(body, &msg)
//		ctx = ctx.WithFromAddressName(msg.Name)
//		ctx = ctx.WithGas(msg.Gas)
//		addr, err := ctx.GetFromAddress()
//		if err != nil {
//			sdk.ErrInvalidAddress("The given Address is Invalid")
//		}
//		ctx = ctx.WithGas(msg.Gas)
//		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
//
//		acc, err := ctx.GetAccountNumber(addr)
//		seq, err := ctx.NextSequence(addr)
//		ctx = ctx.WithSequence(seq)
//		ctx = ctx.WithAccountNumber(acc)
//		if err != nil {
//			w.Write([]byte("account number error"))
//			w.Write([]byte(string(acc)))
//
//		}
//
//		msg1 := sentinel.NewMsgRegisterMasterNode(addr)
//
//		txBytes, err := ctx.sig(msg.Name, msg.Password, []sdk.Msg{msg1}, cdc)
//		if err != nil {
//			w.WriteHeader(http.StatusUnauthorized)
//			w.Write([]byte(err.Error()))
//			return
//		}
//
//		res, err := ctx.BroadcastTx(txBytes)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(err.Error()))
//			return
//		}
//		respon := NewResponse(true, res.Hash.String(), res.Height, res.DeliverTx.Data, res.DeliverTx.Tags)
//		data, err := json.MarshalIndent(respon, "", " ")
//		w.Write(data)
//	}
//	return nil
//}

func validateIp(host string) bool {
	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}