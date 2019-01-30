package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/sentinel"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/gorilla/mux"
)

func ServiceRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {

	r.HandleFunc(
		"/register/vpn", /// service provider
		registervpnHandlerFn(ctx, cdc, kb),
	).Methods("POST")

	//r.HandleFunc(
	//	"/send",
	//	SendTokenHandlerFn(ctx, cdc),
	//).Methods("POST")
	//
	r.HandleFunc(
		"/register/master", // master node
		registermasterdHandlerFn(ctx, cdc, kb),
	).Methods("POST")
	//
	//r.HandleFunc(
	//	"/refund", // client
	//	RefundHandleFn(ctx, cdc),
	//).Methods("POST")
	//
	//r.HandleFunc(
	//	"/master", // owner  or by vote
	//	deleteMasterHandlerFn(ctx, cdc),
	//).Methods("DELETE")
	//
	//r.HandleFunc(
	//	"/vpn", // master node deletes service provider
	//	deleteVpnHandlerFn(ctx, cdc),
	//).Methods("DELETE")
	//r.HandleFunc(
	//	"/vpn/pay", // client
	//	PayVpnServiceHandlerFn(ctx, cdc),
	//).Methods("POST")
	//r.HandleFunc(
	//	"/send-sign", // Off-chain  Tx (client to service provider)
	//	SendSignHandlerFn(),
	//).Methods("POST")
	//r.HandleFunc(
	//	"/vpn/getpayment", // service provider to chain (from kv store)
	//	GetVpnPaymentHandlerFn(ctx, cdc),
	//).Methods("POST")

}

func QueryRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec, keeper sentinel.Keeper) {

	r.HandleFunc(
		"/session/{sessionId}",
		querySessionHandlerFn(cdc, ctx, keeper),
	).Methods("GET")

}

func RegisterRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec, keeper sentinel.Keeper, kb keys.Keybase) {

	ServiceRoutes(ctx, r, cdc, kb)
	QueryRoutes(ctx, r, cdc, keeper)

}
