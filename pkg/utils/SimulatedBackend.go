package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
	Accounts map[common.Address]core.GenesisAccount
	GasLimit uint64
}

func GetSimulatedBackend(c Config) (*backends.SimulatedBackend, error) {
	var backend *backends.SimulatedBackend
	alloc := core.GenesisAlloc(c.Accounts)
	backend = backends.NewSimulatedBackend(alloc, c.GasLimit)
	return backend, nil
}

func GetFaucetAddress() (common.Address, error) {
	var addr common.Address
	faucetSK, err := crypto.GenerateKey()
	if err != nil {
		return addr, err
	}
	addr = crypto.PubkeyToAddress(faucetSK.PublicKey)

	return addr, nil
}
