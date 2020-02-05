package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func getShiftedIndex(index uint32) uint32{
    var e = sdkerrors.ABCIError(ModuleName, index, "")
    if e.Error() == "unknown" {
        return index
    }
    return index + 4;
}

// x/bank module sentinel errors
var (
	ErrNoInputs            = sdkerrors.Register(ModuleName, getShiftedIndex(1), "no inputs to send transaction")
	ErrNoOutputs           = sdkerrors.Register(ModuleName, getShiftedIndex(2), "no outputs to send transaction")
	ErrInputOutputMismatch = sdkerrors.Register(ModuleName, getShiftedIndex(3), "sum inputs != sum outputs")
	ErrSendDisabled        = sdkerrors.Register(ModuleName, getShiftedIndex(4), "send transactions are disabled")
)
