package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sentinel "github.com/cosmos/cosmos-sdk/x/sentinel"
	"github.com/gorilla/mux"
)

func ServiceRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(
		"/register/vpn", /// service provider
		registervpnHandlerFn(ctx, cdc),
	).Methods("POST")
}

func QueryRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec, keeper sentinel.Keeper) {
	r.HandleFunc(
		"/session/{sessionId}",
		querySessionHandlerFn(cdc, ctx, keeper),
	).Methods("GET")
}

func RegisterRoutes(ctx context.CLIContext, r *mux.Router, cdc *codec.Codec, keeper sentinel.Keeper) {
	ServiceRoutes(ctx, r, cdc)
	QueryRoutes(ctx, r, cdc, keeper)
}
