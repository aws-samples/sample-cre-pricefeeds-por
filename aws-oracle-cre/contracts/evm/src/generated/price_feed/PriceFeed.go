// Code generated — DO NOT EDIT.

package price_feed

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var PriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"latestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Structs

// Contract Method Inputs
type OnReportInput struct {
	Metadata []byte
	Report   []byte
}

type SupportsInterfaceInput struct {
	InterfaceId [4]byte
}

type UpdatePriceInput struct {
	Price     *big.Int
	Timestamp *big.Int
}

// Contract Method Outputs
type GetLatestPriceOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
}

type LatestPriceOutput struct {
	Price     *big.Int
	Timestamp *big.Int
}

// Errors

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type PriceUpdatedTopics struct {
}

type PriceUpdatedDecoded struct {
	Price     *big.Int
	Timestamp *big.Int
}

// Main Binding Type for PriceFeed
type PriceFeed struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   PriceFeedCodec
}

type PriceFeedCodec interface {
	EncodeGetLatestPriceMethodCall() ([]byte, error)
	DecodeGetLatestPriceMethodOutput(data []byte) (GetLatestPriceOutput, error)
	EncodeLatestPriceMethodCall() ([]byte, error)
	DecodeLatestPriceMethodOutput(data []byte) (LatestPriceOutput, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	EncodeUpdatePriceMethodCall(in UpdatePriceInput) ([]byte, error)
	PriceUpdatedLogHash() []byte
	EncodePriceUpdatedTopics(evt abi.Event, values []PriceUpdatedTopics) ([]*evm.TopicValues, error)
	DecodePriceUpdated(log *evm.Log) (*PriceUpdatedDecoded, error)
}

func NewPriceFeed(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*PriceFeed, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &PriceFeed{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (PriceFeedCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeGetLatestPriceMethodCall() ([]byte, error) {
	return c.abi.Pack("getLatestPrice")
}

func (c *Codec) DecodeGetLatestPriceMethodOutput(data []byte) (GetLatestPriceOutput, error) {
	vals, err := c.abi.Methods["getLatestPrice"].Outputs.Unpack(data)
	if err != nil {
		return GetLatestPriceOutput{}, err
	}
	if len(vals) != 2 {
		return GetLatestPriceOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetLatestPriceOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetLatestPriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetLatestPriceOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetLatestPriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return GetLatestPriceOutput{
		Arg0: result0,
		Arg1: result1,
	}, nil
}

func (c *Codec) EncodeLatestPriceMethodCall() ([]byte, error) {
	return c.abi.Pack("latestPrice")
}

func (c *Codec) DecodeLatestPriceMethodOutput(data []byte) (LatestPriceOutput, error) {
	vals, err := c.abi.Methods["latestPrice"].Outputs.Unpack(data)
	if err != nil {
		return LatestPriceOutput{}, err
	}
	if len(vals) != 2 {
		return LatestPriceOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return LatestPriceOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return LatestPriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return LatestPriceOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return LatestPriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return LatestPriceOutput{
		Price:     result0,
		Timestamp: result1,
	}, nil
}

func (c *Codec) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Metadata, in.Report)
}

func (c *Codec) EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error) {
	return c.abi.Pack("supportsInterface", in.InterfaceId)
}

func (c *Codec) DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["supportsInterface"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeUpdatePriceMethodCall(in UpdatePriceInput) ([]byte, error) {
	return c.abi.Pack("updatePrice", in.Price, in.Timestamp)
}

func (c *Codec) PriceUpdatedLogHash() []byte {
	return c.abi.Events["PriceUpdated"].ID.Bytes()
}

func (c *Codec) EncodePriceUpdatedTopics(
	evt abi.Event,
	values []PriceUpdatedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodePriceUpdated decodes a log into a PriceUpdated struct.
func (c *Codec) DecodePriceUpdated(log *evm.Log) (*PriceUpdatedDecoded, error) {
	event := new(PriceUpdatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "PriceUpdated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["PriceUpdated"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c PriceFeed) GetLatestPrice(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[GetLatestPriceOutput] {
	calldata, err := c.Codec.EncodeGetLatestPriceMethodCall()
	if err != nil {
		return cre.PromiseFromResult[GetLatestPriceOutput](GetLatestPriceOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetLatestPriceOutput, error) {
		return c.Codec.DecodeGetLatestPriceMethodOutput(response.Data)
	})

}

func (c PriceFeed) LatestPrice(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[LatestPriceOutput] {
	calldata, err := c.Codec.EncodeLatestPriceMethodCall()
	if err != nil {
		return cre.PromiseFromResult[LatestPriceOutput](LatestPriceOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (LatestPriceOutput, error) {
		return c.Codec.DecodeLatestPriceMethodOutput(response.Data)
	})

}

func (c PriceFeed) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

func (c *PriceFeed) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}

// PriceUpdatedTrigger wraps the raw log trigger and provides decoded PriceUpdatedDecoded data
type PriceUpdatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]            // Embed the raw trigger
	contract                        *PriceFeed // Keep reference for decoding
}

// Adapt method that decodes the log into PriceUpdated data
func (t *PriceUpdatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[PriceUpdatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodePriceUpdated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PriceUpdated log: %w", err)
	}

	return &bindings.DecodedLog[PriceUpdatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *PriceFeed) LogTriggerPriceUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []PriceUpdatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[PriceUpdatedDecoded]], error) {
	event := c.ABI.Events["PriceUpdated"]
	topics, err := c.Codec.EncodePriceUpdatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for PriceUpdated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &PriceUpdatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *PriceFeed) FilterLogsPriceUpdated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.PriceUpdatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
