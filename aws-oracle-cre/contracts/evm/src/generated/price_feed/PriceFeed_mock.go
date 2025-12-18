// Code generated — DO NOT EDIT.

//go:build !wasip1

package price_feed

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

// PriceFeedMock is a mock implementation of PriceFeed for testing.
type PriceFeedMock struct {
	GetLatestPrice func() (GetLatestPriceOutput, error)
	LatestPrice    func() (LatestPriceOutput, error)
}

// NewPriceFeedMock creates a new PriceFeedMock for testing.
func NewPriceFeedMock(address common.Address, clientMock *evmmock.ClientCapability) *PriceFeedMock {
	mock := &PriceFeedMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["getLatestPrice"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetLatestPrice == nil {
				return nil, errors.New("getLatestPrice method not mocked")
			}
			result, err := mock.GetLatestPrice()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getLatestPrice"].Outputs.Pack(
				result.Arg0,
				result.Arg1,
			)
		},
		string(abi.Methods["latestPrice"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.LatestPrice == nil {
				return nil, errors.New("latestPrice method not mocked")
			}
			result, err := mock.LatestPrice()
			if err != nil {
				return nil, err
			}
			return abi.Methods["latestPrice"].Outputs.Pack(
				result.Price,
				result.Timestamp,
			)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
