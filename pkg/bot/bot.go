package bot

import (
	// "context"
	"math/big"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/dao"
)

// GetBalance Get the balance of the DOA's Pool
func GetBalance(d dao.DAO) (*big.Int, error) {

	var balance *big.Int
	// client := d.Instance
	//
	// balance, err := client.BalanceAt(context.Background(), d.Address, nil)
	// if err != nil {
	// 	return balance, err
	// }
	//
	return balance, nil
}
