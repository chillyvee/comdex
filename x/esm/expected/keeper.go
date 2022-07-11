package expected

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AssetKeeper interface {
	GetApp(ctx sdk.Context, id uint64) (app assettypes.AppData, found bool)
	GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool)
	GetApps(ctx sdk.Context) (apps []assettypes.AppData, found bool)
	GetAssetsForOracle(ctx sdk.Context) (assets []assettypes.Asset)
}

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, name string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, name string, address sdk.AccAddress, coins sdk.Coins) error
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(
		ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins,
	) error
	SendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type MarketKeeper interface {
	GetPriceForAsset(ctx sdk.Context, id uint64) (uint64, bool)
}