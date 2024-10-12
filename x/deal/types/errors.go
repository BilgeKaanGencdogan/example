package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/deal module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
)
var (
	ErrInvalidOwner           = sdkerrors.Register(ModuleName, 1102, "Owner does not exists")
	ErrInvalidVendor          = sdkerrors.Register(ModuleName, 1103, "Vendor does not exists")
	ErrInvalidConsumer        = sdkerrors.Register(ModuleName, 1104, "Consumer does not exists")
	ErrInvalidCommission      = sdkerrors.Register(ModuleName, 1105, "Invalid commission")
	ErrDealNotFound           = sdkerrors.Register(ModuleName, 1106, "Deal not found")
	ErrContractNotFound       = sdkerrors.Register(ModuleName, 1107, "Contract not found")
	ErrContractExpired        = sdkerrors.Register(ModuleName, 1108, "Contract already expired")
	ErrInvalidTime            = sdkerrors.Register(ModuleName, 1109, "Invalid Time")
	ErrDescLength             = sdkerrors.Register(ModuleName, 1110, "Desc length too large")
	ErrPaymentFailed          = sdkerrors.Register(ModuleName, 1111, "Payment Failed")
	ErrNotCommitted           = sdkerrors.Register(ModuleName, 1112, "Contract not in commit stage")
	ErrNotShipped             = sdkerrors.Register(ModuleName, 1113, "Order not yet shipped or completed")
	ErrNotApprovedOrCompleted = sdkerrors.Register(ModuleName, 1114, "Contract is either completed or not in approved stage")
	ErrRefund                 = sdkerrors.Register(ModuleName, 1115, "Refund applicable after 20 mins delay")
	ErrVendorETA              = sdkerrors.Register(ModuleName, 1116, "Vendor ETA should be less than or equal to half of that of Owner ETA")
	ErrInvalidAddress         = sdkerrors.Register(ModuleName, 1117, "invalid creator address")
)
