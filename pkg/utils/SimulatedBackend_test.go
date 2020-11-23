package utils

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	faucetAddr, err := GetFaucetAddress()
	if err != nil {
		t.Error(err)
	}

	balAddr0 := big.NewInt(1)
	balFaucet := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: balAddr0}, // ECRecover
		faucetAddr:                       {Balance: balFaucet},
	}

	config := Config{addr, 9000000}
	underTest, err := GetSimulatedBackend(config)
	if err != nil {
		t.Error(err)
	}

	// balance of each address created
	actual, err := underTest.BalanceAt(context.Background(), common.BytesToAddress([]byte{1}), big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, balAddr0, actual)

	actual, err = underTest.BalanceAt(context.Background(), faucetAddr, big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, balFaucet, actual)
}
