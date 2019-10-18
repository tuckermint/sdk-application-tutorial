# Codec File

To [register your types with Amino](https://github.com/tendermint/go-amino#registering-types) so that they can be encoded/decoded, there is a bit of code that needs to be placed in `./x/tuckermint/types/codec.go`. Any interface you create and any struct that implements an interface needs to be declared in the `RegisterCodec` function. In this module the three `Msg` implementations (`SetName`, `BuyName` and `DeleteName`) need to be registered, but your `Whois` query return type does not. In addition, we define a module specific codec for use later.

```go
package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "tuckermint/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "tuckermint/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "tuckermint/DeleteName", nil)
}
```

### Next you need to define [CLI interactions](./cli.md) with your module.
