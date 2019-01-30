package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type ClientSignature struct {
	Coins     sdk.Coins
	Sessionid []byte
	Counter   int64
	Signature Signature
	IsFinal   bool
}
type Signature struct {
	Pubkey    crypto.PubKey
}

func NewClientSignature(coins sdk.Coins, sesid []byte, counter int64, pubkey crypto.PubKey,  isfinal bool) ClientSignature {
	return ClientSignature{
		Coins:     coins,
		Sessionid: sesid,
		Counter:   counter,
		IsFinal:   isfinal,
		Signature: Signature{
			Pubkey:    pubkey,
		},
	}
}
func (a ClientSignature) Value() Signature {
	return a.Signature
}

type StdSig struct {
	Coins     sdk.Coins
	Sessionid []byte
	Counter   int64
	Isfinal   bool
}

func ClientStdSignBytes(coins sdk.Coins, sessionid []byte, counter int64, isfinal bool) []byte {
	bz, err := json.Marshal(StdSig{
		Coins:     coins,
		Sessionid: sessionid,
		Counter:   counter,
		Isfinal:   isfinal,
	})
	if err != nil {
	}
	return sdk.MustSortJSON(bz)
}

type Vpnsign struct {
	From     sdk.AccAddress
	Ip       string
	Netspeed int64
	Ppgb     int64
	Location string
}

func GetVPNSignature(address sdk.AccAddress, ip string, ppgb int64, netspeed int64, location string) []byte {
	bz, err := json.Marshal(Vpnsign{
		From:     address,
		Ip:       ip,
		Ppgb:     ppgb,
		Netspeed: netspeed,
		Location: location,
	})
	if err != nil {

	}
	return sdk.MustSortJSON(bz)
}