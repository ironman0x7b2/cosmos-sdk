package rest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/sentinel"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	ioutill "io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name Account name of service provider.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
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
func registervpnHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
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
			fmt.Println(err)
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

		//var req MsgRegisterVpnService
		//err = utils.ReadRESTReq(w, r, cdc, &req)

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()

		msg1 := sentinel.NewMsgRegisterVpnService(addr, msg.Ip, msg.Ppgb, msg.UploadSpeed, msg.DownloadSpeed, msg.EncMethod, msg.Latitude, msg.Longitude, msg.City, msg.Country, msg.NodeType, msg.Version)
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)

	}
	return nil
}


func registermasterdHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := MsgRegisterMasterNode{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}

		json.Unmarshal(body, &msg)

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()


		msg1 := sentinel.NewMsgRegisterMasterNode(addr)

		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)

	}
	return nil
}

/**
* @api {delete} /vpn To Delete VPN Node.
* @apiName  deleteVpnNode
* @apiGroup Sentinel-Tendermint
* @apiParam {String} address  Address of VPN Node which we want to delete.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the person who is deleting the VPN node.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
* @apiError AccountNotExists VPN Node not exists
* @apiErrorExample AccountNotExists-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      13
* ABCICode:  1245197
* Error:     --= Error =--
* Data: common.FmtError{format:"Account is not exist", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiSuccessExample Response:
{
 *   "Success": true,
 *   "Hash": "32EF9DFB6BC24D3159A8310F1AE438BED479466E",
 *   "Height": 3698,
 *   "Data": "FRTjZrQKAswn4UTeyJ0eXBlIjoic2VudGluZWWQiOiIxMiIsIlBwZ2IiOiIyMyIsIkxvY2F0aW9uIjoiaHlkIn19b1W/Usl/KB3iflg==",
 *   "Tags": [
 *       {
 *           "key": "ZGVsZXRlZCBWcG4gYWRkcmVzcw==",
 *           "value": "42a0CgLMJ+FE29Vv1LJfygd4n5Y="
 *      }
 *  ]
}
*/

func deleteVpnHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var msg MsgDeleteVpnUser
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		json.Unmarshal(body, &msg)
		if msg.Address == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid address."))
			return
		}
		Vaddr, err := sdk.AccAddressFromBech32(msg.Address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()
		msg1 := sentinel.NewMsgDeleteVpnUser(addr, Vaddr)
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)

	}
	return nil
}

func deleteMasterHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var msg MsgDeleteMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		json.Unmarshal(body, &msg)

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}

	}
	return nil
}

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
//
/**
* @api {post} /vpn/pay To Pay for VPN service.
* @apiName  payVPN service
* @apiGroup Sentinel-Tendermint
* @apiParam {String} amount  Amount to pay for vpn service.
* @apiParam {String} vaddress Address of the vpn service provider.
* @apiParam {String} sig_name NewAccountName.
* @apiParam {String} sig_password NewAccountPassword.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the client.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
* @apiError AccountNotExists VPN Node not exists
* @apiError AccountNameAlreadyExists The new account name is already exist
* @apiError InsufficientFunds  Insufficient Funds
* @apiErrorExample AccountVPNNotExists-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 1
* Code:      9
* ABCICode:  65545
* Error:     --= Error =--
* Data: common.FmtError{format:"VPN address is not registered", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiErrorExample AccountNameAlreadyExists-Response:
*{
	* " Account with name XXXXXX already exists."
*}
* @apiErrorExample InsufficientFunds-Response:
*{
	* "Insufficient Funds"
*}
* @apiSuccessExample Response:
{
*   "Success": true,
*   "Hash": "D2C58CAFC580CC39A4CFAB4325991A9378AFE77D",
*   "Height": 1196,
*   "Data": "IjNwWGdHazB5MnBGceyJ0eXBlIjoic2VudGluZWwvcmVnaXN0ZXJ2cG4iLCJ2YWx1ZSI6eyJGc3BlZWQiOiIxMiIsIlBwZ2IiOiIyMyIsIkxvY2F0aW9uIjoiaHlkIn19TdZdWIwak5xIg==",
*   "Tags": [
*      {
*       "key": "c2VuZGVyIGFkZHJlc3M=",
*       "value": "Y29zbW9zYWNjYWRkcjFuY3hlbGpjcjRnOWhzdmw3amRuempkazNyNzYyamUzenk4bXU5MA=="
*      },
*     {
*      "key": "c2Vlc2lvbiBpZA==",
*      "value": "M3BYZ0drMHkycEZxN1l1YjBqTnE="
*     }
*          ]
}
*/
func PayVpnServiceHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		msg := MsgPayVpnService{}
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)

		seed := getSeed(keys.Secp256k1)
		if msg.SigName == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Enter the Name."))
			return
		}
		if msg.SigPassword == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Enter the Password."))
			return
		}

		if msg.Coins == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Enter the coins "))
			return
		}
		if msg.Vpnaddr == "" {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid vpn address"))
			return
		}
		vaddr, err := sdk.AccAddressFromBech32(msg.Vpnaddr)
		fmt.Println("vaddr:",vaddr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		coins, err := sdk.ParseCoins(msg.Coins)
		if err != nil {

			sdk.ErrInternal("Parse Coins Failed")
		}

		infos, err := kb.List()
		for _, i := range infos {
			if i.GetName() == msg.SigName {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte(fmt.Sprintf("Account with name %s already exists.", msg.SigName)))
				return
			}
		}
		info, err := kb.CreateKey(msg.SigName, seed, msg.SigPassword)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		keyinfo, err := kb.Get(msg.BaseReq.Name)
		addr := sdk.AccAddress(keyinfo.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		if err != nil {
			fmt.Println("Error while getting res: ",err)
		}
		fmt.Println("res is: ",res)
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()
		msg1 := sentinel.NewMsgPayVpnService(coins, vaddr, addr, info.GetPubKey())
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)

	}
	return nil
}

//
//To create client signature....... This is not a transaction......

/**
* @api {post} /send-sign To Create sigature of the client.
* @apiName  CreateSignature
* @apiGroup Sentinel-Tendermint
* @apiParam {String} session_id session-id.
* @apiParam {String} amount Amount to create signature.
* @apiParam {Number} counter Counter value of the sigature.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the client.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
*@apiParam {Boolean} isfial boolean value for is this final signature or not.
* @apiSuccessExample Response:
* 10lz2f928xpzsyggqhc9mu80qj59vx0rc6sedxmsfhca8ysuhhtgqypar3h4ty0pgftwqygp6vm54drttw5grlz4p5n238cvzxe2vpxmu6hhnqvt0uxstg7et4vdqhm4v
*/

func SendSignHandlerFn(kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := ClientSignature{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		}
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		if msg.BaseReq.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid Account Name."))
			return
		}
		if msg.BaseReq.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid Password."))
			return
		}
		coins, err := sdk.ParseCoins(msg.Coins)
		if err != nil {
			sdk.ErrInternal("Parse Coins Failed")
		}

		bz := clientStdSignBytes(coins, []byte(msg.Sessionid), msg.Counter, msg.IsFinal)
		
		fmt.Println("}}}}}}}}}}}}}}}}}}}}}}}}}}]",string(bz))

		sign, _, err := kb.Sign(msg.BaseReq.Name, msg.BaseReq.Password, bz)
		if err != nil {
			w.Write([]byte(" Signature failed"))
			return
		}
		 w.Write([]byte(base64.StdEncoding.EncodeToString(sign)))
	}
	return nil
}

type stdSig struct {
	Coins     sdk.Coins
	Sessionid []byte
	Counter   int64
	Isfinal   bool
}

func clientStdSignBytes(coins sdk.Coins, sessionid []byte, counter int64, isfinal bool) []byte {
	bz, err := json.Marshal(stdSig{
		Coins:     coins,
		Sessionid: sessionid,
		Counter:   counter,
		Isfinal:   isfinal,
	})
	if err != nil {
	}
	return sdk.MustSortJSON(bz)
}

/**
* @api {post} /refund To Refund the balance of client.
* @apiName  Refund
* @apiGroup Sentinel-Tendermint
* @apiParam {String} session_id session-id.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the client.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
* @apiError TimeInvalidError Time is not more than 24 hours
* @apiError InvalidSessionIdError SessionId is invalid
* @apiErrorExample TimeInvalidError-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      2
* ABCICode:  6551245
* Error:     --= Error =--
* Data: common.FmtError{format:"time is less than 24 hours  or the balance is negative or equal to zero", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiErrorExample InvalidSessionIdError-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      6
* ABCICode:  124545
* Error:     --= Error =--
* Data: common.FmtError{format:"Invalid SessionId", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiSuccessExample Response:
*{
*	{
*   "Success": true,
*   "Hash": "868B602828FA48F1D4A03D9D066EB42DEC483AA0",
*   "Height": 1092,
*   "Data": "Qwi/dQ1h0GcdrppVOeyJ0eXBlIjoic2VudGluZWwvcmVnaXN0yJGc3BlZWQiOiIxMiIsIlBwZ2IiOiIyMyIsIkxvY2F0aW9uIjoiaHlkIn192hhGfJVl3g=",
*   "Tags": [
*{
*           "key": "Y2xpZW50IFJlZnVuZCBBZGRyZXNzOg==",
*           "value": "Y29zbW9zYWNjYWRkcjFndnl0N2FnZHY4Z3h3OGR3bmYybms2cnByOGU5dDltY3hkeGV3cA=="
*       }
*   ]
*}
* }
*/

func RefundHandleFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := MsgRefund{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			sentinel.ErrUnMarshal("Unmarshal of Given Message Type is failed")

		}

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()

		msg1 := sentinel.NewMsgRefund(addr, []byte(msg.Sessionid))
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)
	}
}

/**
* @api {post} /vpn/getpayment To get payment of vpn service
* @apiName  GetVPNPayment
* @apiGroup Sentinel-Tendermint
* @apiParam {String} amount Amount to send VPN node.
* @apiParam {String} session_id session-id.
* @apiParam {Number} counter Counter value.
* @apiParam {Boolean} isfinal is this final signature or not.
* @apiParam {string} sign signature of the client.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the vpn service provider.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
* @apiError InvalidSessionId  SessionId is invalid
* @apiError SignatureVerificationFailed  Invalid signature
* @apiErrorExample InvalidSessionId-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      6
* ABCICode:  65545
* Error:     --= Error =--
* Data: common.FmtError{format:"Invalid session Id", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiErrorExample SignatureVerificationFailed-Response:
*{
* checkTx failed: (1245197) Msg 0 failed: === ABCI Log ===
* Codespace: 19
* Code:      6
* ABCICode:  65545
* Error:     --= Error =--
* Data: common.FmtError{format:"signature verification failed", args:[]interface {}(nil)}
* Msg Traces:
* --= /Error =--
*
*=== /ABCI Log ===
*}
* @apiSuccessExample Response:
*{
*    "Success": true,
*    "Hash": "629F4603A5A4DE598B58DC494CCC38DB9FD96604",
*    "Height": 353,
*    "Data":"eyJ0eXBlIjoic2VudGluZWwvcmVnaXN0ZXJ2cG4iLCJ2YWx1ZSI6eyJGc3BlZWQiOiIxMiIsIlBwZ2IiOiyJ0eXBlIjoic2VudGluZWwvcmVnaXN0ZXJ2cG4iLCJ2YW9==",
*    "Tags": [
*        {
*            "key": "VnBuIFByb3ZpZGVyIEFkZHJlc3M6",
*            "value": "Y29zbW9zYWNjYWRkcjF1ZG50Z3pzemVzbjd6M3htNjRoYWZ2amxlZ3JoMzh1a3p3OW03Zw=="
*        },
*        {
*            "key": "c2Vlc2lvbklk",
*            "value": "WVZJRW81Y0dIczdkb09UVzRDTk4="
*        }
*    ]
*}
*/

func GetVpnPaymentHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := MsgGetVpnPayment{}
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			sentinel.ErrUnMarshal("UnMarshal of MessageType is failed")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if msg.Signature == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid signature ."))
			return
		}
		if msg.Coins == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid amount."))
			return
		}
		if msg.Sessionid == "" {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Session Id is wrong"))
			return
		}
		if msg.Counter < 0 {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Counter"))
			return
		}
		coins, err := sdk.ParseCoins(msg.Coins)
		if err != nil {
			sdk.ErrInternal("Parse Coins failed")
		}

		sig , err := base64.StdEncoding.DecodeString(msg.Signature)
		if err!=nil{
			// Handle your error
		}

		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		if err != nil {
			sdk.ErrInvalidAddress("The given Address is Invalid")
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()
		msg1 := sentinel.NewMsgGetVpnPayment(coins, []byte(msg.Sessionid), msg.Counter,  sig,addr, msg.IsFinal)

		fmt.Println("Message is ", msg1)
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)
	}
	return nil
}

func getBech64Signature(address string) (pkk crypto.PubKey, err error) {
	hrp, bz, err := Decode(address)
	fmt.Println("hrp is: ",hrp)
	fmt.Println("bz is: ",bz)
	fmt.Println("err in decoding is: ",err)
	if err != nil {
		return nil, err
	}
	if hrp != "" {
		return nil, fmt.Errorf("invalid bech32 prefix. Expected %s, Got %s", "", hrp)
	}

	//pk, err = cryptoAmino.SignatureFromBytes(bz)  
	// pk, err = cryptoAmino.PubKeyFromBytes(bz)
	pk, err := pubKeyFromBytes(bz)
	fmt.Println("pk is: ",pk)
	fmt.Println("err in pk is: ",pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}


func pubKeyFromBytes(pubKeyBytes []byte) (pubKey crypto.PubKey, err error) {
	err = cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	fmt.Println("err while getting pubkeyfromBytes: ",err)
	return
}

/**
* @api {post} /send To send money to account.
* @apiName sendTokens
* @apiGroup Sentinel-Tendermint
* @apiParam {String} to To address.
* @apiParam {String} amount Amount to send.
* @apiParam {Object} base_req Base request Object
* @apiParam {String} base_req.name AccountName of the sender.
* @apiParam {string} base_req.password Password of account.
* @apiParam {string} base_req.chain_id Chain ID
*
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

func SendTokenHandlerFn(ctx context.CLIContext, cdc *codec.Codec, kb keys.Keybase, decoder auth.AccountDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := SendTokens{}
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			sentinel.ErrUnMarshal("UnMarshal of MessageType is failed")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if msg.Coins == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid amount."))
			return
		}
		to, err := sdk.AccAddressFromBech32(msg.ToAddress)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		coins, err := sdk.ParseCoins(msg.Coins)
		if err != nil {
			sdk.ErrInternal("Parse Coins failed")
		}
		info, err := kb.Get(msg.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr := sdk.AccAddress(info.GetPubKey().Address())
		fmt.Println("Address is: ",addr)
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), "acc")
		account, err := decoder(res)
		if err != nil {
			fmt.Println("Error while decoding account: ",account)
		}

		msg.BaseReq.AccountNumber = account.GetAccountNumber()
		msg.BaseReq.Sequence = account.GetSequence()
		msg1 := sentinel.NewMsgSendTokens(addr, coins, to)
		utils.CompleteAndBroadcastTxREST(w, r, ctx, msg.BaseReq, []sdk.Msg{msg1}, cdc)
	}
}
