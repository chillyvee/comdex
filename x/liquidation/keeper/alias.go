package keeper

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	auctiontypes "github.com/comdex-official/comdex/x/auction/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k *Keeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	return k.account.GetModuleAccount(ctx, name)
}

func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k.bank.GetBalance(ctx, addr, denom)
}

func (k *Keeper) MintCoin(ctx sdk.Context, name string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.MintCoins(ctx, name, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}

func (k *Keeper) GetVaults(ctx sdk.Context) (vaults []vaulttypes.Vault) {
	return k.vault.GetVaults(ctx)
}

func (k *Keeper) DeleteVault(ctx sdk.Context, id uint64) {
	k.vault.DeleteVault(ctx, id)
}

func (k *Keeper) DeleteVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, pairID uint64) {
	k.vault.DeleteVaultForAddressByPair(ctx, address, pairID)
}

func (k *Keeper) GetPair(ctx sdk.Context, id uint64) (assettypes.Pair, bool) {
	return k.asset.GetPair(ctx, id)
}

func (k *Keeper) GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool) {
	return k.asset.GetAsset(ctx, id)
}

func (k *Keeper) CalculateCollaterlizationRatio(
	ctx sdk.Context,
	amountIn sdk.Int,
	assetIn assettypes.Asset,
	amountOut sdk.Int,
	assetOut assettypes.Asset,
) (sdk.Dec, error) {
	return k.vault.CalculateCollaterlizationRatio(ctx, amountIn, assetIn, amountOut, assetOut)
}

func (k *Keeper) CreteNewVault(ctx sdk.Context, pairdId uint64, from string, assetIn assettypes.Asset, amountIn sdk.Int, assetOut assettypes.Asset, amountOut sdk.Int) error {
	return k.vault.CreteNewVault(ctx, pairdId, from, assetIn, amountIn, assetOut, amountOut)
}

func (k *Keeper) GetVaultID(ctx sdk.Context) uint64 {
	return k.vault.GetID(ctx)
}

func (k *Keeper) SetVaultID(ctx sdk.Context, id uint64) {
	k.vault.SetID(ctx, id)
}

func (k *Keeper) SetVault(ctx sdk.Context, vault vaulttypes.Vault) {
	k.vault.SetVault(ctx, vault)
}

func (k *Keeper) SetVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, pairID uint64, id uint64) {
	k.vault.SetVaultForAddressByPair(ctx, address, pairID, id)
}

func (k *Keeper) SendCoinFromModuleToAccount(ctx sdk.Context, name string, address sdk.AccAddress, coin sdk.Coin) error {
	return k.bank.SendCoinsFromModuleToAccount(ctx, name, address, sdk.NewCoins(coin))
}

func (k *Keeper) GetPriceForAsset(ctx sdk.Context, id uint64) (uint64, bool) {
	return k.market.GetPriceForAsset(ctx, id)
}

func (k *Keeper) UpdateUserVaultIdMapping(ctx sdk.Context, vaultOwner string, vaultId uint64, isInsert bool) error {
	return k.vault.UpdateUserVaultIdMapping(ctx, vaultOwner, vaultId, isInsert)
}

func (k *Keeper) UpdateCollateralVaultIdMapping(ctx sdk.Context, assetInDenom string, assetOutDenom string, vaultId uint64, isInsert bool) error {
	return k.vault.UpdateCollateralVaultIdMapping(ctx, assetInDenom, assetOutDenom, vaultId, isInsert)
}

func (k *Keeper) GetAuctionParams(ctx sdk.Context) auctiontypes.Params {
	return k.auction.GetParams(ctx)
}
