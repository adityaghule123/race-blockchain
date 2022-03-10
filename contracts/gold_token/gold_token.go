package gold_token

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/contracts"
	"github.com/celo-org/celo-blockchain/contracts/abis"
	"github.com/celo-org/celo-blockchain/core/vm"
	"github.com/celo-org/celo-blockchain/params"
)

var (
	totalSupplyMethod    = contracts.NewRegisteredContractMethod(params.RaceTokenRegistryId, abis.RaceToken, "totalSupply", params.MaxGasForTotalSupply)
	increaseSupplyMethod = contracts.NewRegisteredContractMethod(params.RaceTokenRegistryId, abis.RaceToken, "increaseSupply", params.MaxGasForIncreaseSupply)
	mintMethod           = contracts.NewRegisteredContractMethod(params.RaceTokenRegistryId, abis.RaceToken, "mint", params.MaxGasForMintGas)
)

func GetTotalSupply(vmRunner vm.EVMRunner) (*big.Int, error) {
	var totalSupply *big.Int
	err := totalSupplyMethod.Query(vmRunner, &totalSupply)
	return totalSupply, err
}

func IncreaseSupply(vmRunner vm.EVMRunner, value *big.Int) error {
	err := increaseSupplyMethod.Execute(vmRunner, nil, common.Big0, value)
	return err
}

func Mint(vmRunner vm.EVMRunner, beneficiary common.Address, value *big.Int) error {
	if value.Cmp(new(big.Int)) <= 0 {
		return nil
	}

	err := mintMethod.Execute(vmRunner, nil, common.Big0, beneficiary, value)
	return err
}
