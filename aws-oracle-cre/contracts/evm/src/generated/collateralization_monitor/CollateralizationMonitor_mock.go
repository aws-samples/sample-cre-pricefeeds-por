// Code generated — DO NOT EDIT.

//go:build !wasip1

package collateralization_monitor

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// CollateralizationMonitorMock is a mock implementation of CollateralizationMonitor for testing.
type CollateralizationMonitorMock struct {
	GetLatestData func() (CollateralData, error)
	LatestData    func() (LatestDataOutput, error)
	MinRatio      func() (*big.Int, error)
}

// NewCollateralizationMonitorMock creates a new CollateralizationMonitorMock for testing.
func NewCollateralizationMonitorMock(address common.Address, clientMock *evmmock.ClientCapability) *CollateralizationMonitorMock {
	mock := &CollateralizationMonitorMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["getLatestData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetLatestData == nil {
				return nil, errors.New("getLatestData method not mocked")
			}
			result, err := mock.GetLatestData()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getLatestData"].Outputs.Pack(result)
		},
		string(abi.Methods["latestData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.LatestData == nil {
				return nil, errors.New("latestData method not mocked")
			}
			result, err := mock.LatestData()
			if err != nil {
				return nil, err
			}
			return abi.Methods["latestData"].Outputs.Pack(
				result.Price,
				result.Reserves,
				result.Ratio,
				result.Timestamp,
				result.IsHealthy,
			)
		},
		string(abi.Methods["minRatio"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MinRatio == nil {
				return nil, errors.New("minRatio method not mocked")
			}
			result, err := mock.MinRatio()
			if err != nil {
				return nil, err
			}
			return abi.Methods["minRatio"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
