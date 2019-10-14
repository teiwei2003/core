// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-project/core/x/market/internal/types/
// ALIASGEN: github.com/terra-project/core/x/market/internal/keeper/
package market

import (
	"github.com/terra-project/core/x/market/internal/keeper"
	"github.com/terra-project/core/x/market/internal/types"
)

const (
	DefaultCodespace      = types.DefaultCodespace
	CodeInsufficientSwap  = types.CodeInsufficientSwap
	CodeNoEffectivePrice  = types.CodeNoEffectivePrice
	CodeRecursiveSwap     = types.CodeRecursiveSwap
	CodeInactive          = types.CodeInactive
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	RouterKey             = types.RouterKey
	QuerierRoute          = types.QuerierRoute
	DefaultParamspace     = types.DefaultParamspace
	QuerySwap             = types.QuerySwap
	QueryTerraPool        = types.QueryTerraPool
	QueryBasePool         = types.QueryBasePool
	QueryLastUpdateHeight = types.QueryLastUpdateHeight
	QueryParameters       = types.QueryParameters
)

var (
	// functions aliases
	RegisterCodec            = types.RegisterCodec
	ErrNoEffectivePrice      = types.ErrNoEffectivePrice
	ErrInsufficientSwapCoins = types.ErrInsufficientSwapCoins
	ErrRecursiveSwap         = types.ErrRecursiveSwap
	ErrInactive              = types.ErrInactive
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	NewMsgSwap               = types.NewMsgSwap
	DefaultParams            = types.DefaultParams
	NewQuerySwapParams       = types.NewQuerySwapParams
	NewKeeper                = keeper.NewKeeper
	ParamKeyTable            = keeper.ParamKeyTable
	NewQuerier               = keeper.NewQuerier

	// variable aliases
	ModuleCdc                             = types.ModuleCdc
	BasePoolKey                           = types.BasePoolKey
	TerraPoolKey                          = types.TerraPoolKey
	LastUpdateHeightKey                   = types.LastUpdateHeightKey
	ParamStoreKeyPoolUpdateInterval       = types.ParamStoreKeyPoolUpdateInterval
	ParamStoreKeyTerraLiquidityRatio = types.ParamStoreKeyTerraLiquidityRatio
	ParamStoreKeyMinSpread                = types.ParamStoreKeyMinSpread
	ParmamStoreKeyTobinTax                = types.ParmamStoreKeyTobinTax
	DefaultPoolUpdateInterval             = types.DefaultPoolUpdateInterval
	DefaultTerraLiquidityRatio       = types.DefaultTerraLiquidityRatio
	DefaultMinSpread                      = types.DefaultMinSpread
	DefaultTobinTax                       = types.DefaultTobinTax
)

type (
	SupplyKeeper    = types.SupplyKeeper
	OracleKeeper    = types.OracleKeeper
	GenesisState    = types.GenesisState
	MsgSwap         = types.MsgSwap
	Params          = types.Params
	QuerySwapParams = types.QuerySwapParams
	Keeper          = keeper.Keeper
)
