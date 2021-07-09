package banking

import (
	"errors"
	"math/big"

	"code.vegaprotocol.io/data-node/assets"
	"code.vegaprotocol.io/data-node/assets/common"
	"code.vegaprotocol.io/data-node/types"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrUnknownAssetAction = errors.New("unknown asset action")
)

type withdrawal struct {
	nonce *big.Int
}

type txRef struct {
	asset       common.AssetClass
	blockNumber uint64
	hash        string
	logIndex    uint64
}

type assetAction struct {
	id    string
	state uint32
	asset *assets.Asset

	// erc20 specifics
	blockNumber uint64
	txIndex     uint64
	hash        string

	// all deposit related types
	builtinD *types.BuiltinAssetDeposit
	erc20D   *types.ERC20Deposit

	// all asset list related types
	withdrawal *withdrawal
	erc20AL    *types.ERC20AssetList

	// all withdrawal related types
	erc20W *types.ERC20Withdrawal
}

func (t *assetAction) GetID() string {
	return t.id
}

func (t *assetAction) IsBuiltinAssetDeposit() bool {
	return t.builtinD != nil
}

func (t *assetAction) IsERC20Deposit() bool {
	return t.erc20D != nil
}

func (t *assetAction) IsERC20Withdrawal() bool {
	return t.erc20W != nil
}

func (t *assetAction) IsERC20AssetList() bool {
	return t.erc20AL != nil
}

func (t *assetAction) BuiltinAssetDesposit() *types.BuiltinAssetDeposit {
	return t.builtinD
}

func (t *assetAction) ERC20Deposit() *types.ERC20Deposit {
	return t.erc20D
}

func (t *assetAction) ERC20Withdrawal() *types.ERC20Withdrawal {
	return t.erc20W
}

func (t *assetAction) ERC20AssetList() *types.ERC20AssetList {
	return t.erc20AL
}

func (t *assetAction) String() string {
	switch {
	case t.IsBuiltinAssetDeposit():
		return t.builtinD.String()
	case t.IsERC20Deposit():
		return t.erc20D.String()
	case t.IsERC20AssetList():
		return t.erc20AL.String()
	case t.IsERC20Withdrawal():
		return t.erc20W.String()
	default:
		return ""
	}
}

func (t *assetAction) Check() error {
	switch {
	case t.IsBuiltinAssetDeposit():
		return t.checkBuiltinAssetDeposit()
	default:
		return ErrUnknownAssetAction
	}
}

func (t *assetAction) checkBuiltinAssetDeposit() error {
	return nil
}

func (t *assetAction) getRef() txRef {
	switch {
	case t.IsBuiltinAssetDeposit():
		return txRef{common.Builtin, 0, uuid.NewV4().String(), 0}
	case t.IsERC20Deposit():
		return txRef{common.ERC20, t.blockNumber, t.hash, t.txIndex}
	case t.IsERC20AssetList():
		return txRef{common.ERC20, t.blockNumber, t.hash, t.txIndex}
	case t.IsERC20Withdrawal():
		return txRef{common.ERC20, t.blockNumber, t.hash, t.txIndex}
	default:
		return txRef{} // this is basically unreachable
	}
}
