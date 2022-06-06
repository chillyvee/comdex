package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	collectorKeeper "github.com/comdex-official/comdex/x/collector/keeper"
	liquidationKeeper "github.com/comdex-official/comdex/x/liquidation/keeper"
	lockerkeeper "github.com/comdex-official/comdex/x/locker/keeper"
	rewardsKeeper "github.com/comdex-official/comdex/x/rewards/keeper"
	tokenMintkeeper "github.com/comdex-official/comdex/x/tokenmint/keeper"
)

func RegisterCustomPlugins(
	locker *lockerkeeper.Keeper,
	tokenMint *tokenMintkeeper.Keeper,
	asset *assetkeeper.Keeper,
	rewards *rewardsKeeper.Keeper,
	collector *collectorKeeper.Keeper,
	liquidation *liquidationKeeper.Keeper,
) []wasmkeeper.Option {

	comdexQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation)

	appDataqueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(comdexQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation),
	)

	return []wasm.Option{
		appDataqueryPluginOpt,
		messengerDecoratorOpt,
	}
}