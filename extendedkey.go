package secp256k1

import (
	"encoding/json"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

const (
	RecommendedSeedLen = hdkeychain.RecommendedSeedLen
	MinSeedBytes       = hdkeychain.MinSeedBytes
	MaxSeedBytes       = hdkeychain.MaxSeedBytes
	HardenedKeyStart   = hdkeychain.HardenedKeyStart
)

type ExtendedKey struct {
	hdkeychain.ExtendedKey
}

func NewMaster(seed []byte, net *chaincfg.Params) (*ExtendedKey, error) {
	k, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return nil, err
	}
	return &ExtendedKey{*k}, nil
}

func NewKeyFromString(key string) (*ExtendedKey, error) {
	k, err := hdkeychain.NewKeyFromString(key)
	if err != nil {
		return nil, err
	}
	return &ExtendedKey{*k}, nil
}

func (k *ExtendedKey) Child(i uint32) (*ExtendedKey, error) {
	key, err := k.ExtendedKey.Child(i)
	if err != nil {
		return nil, err
	}
	return &ExtendedKey{*key}, nil
}

func (k *ExtendedKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}

func (k *ExtendedKey) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	key, err := NewKeyFromString(s)
	if err != nil {
		return err
	}
	*k = *key
	return nil
}

func GenerateSeed(length uint8) ([]byte, error) {
	return hdkeychain.GenerateSeed(length)
}
