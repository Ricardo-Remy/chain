package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

func GetWasmOpts(appOpts servertypes.AppOptions) []wasm.Option {
	var wasmOpts []wasm.Option
	return wasmOpts
}
