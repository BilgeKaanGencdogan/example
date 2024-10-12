package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateContract = "create_contract"

var _ sdk.Msg = &MsgCreateContract{}

func NewMsgCreateContract(creator string, dealId string, consumer string, desc string, ownerETA string, expiry string, fees uint64) *MsgCreateContract {
	return &MsgCreateContract{
		Creator:  creator,
		DealId:   dealId,
		Consumer: consumer,
		Desc:     desc,
		OwnerETA: ownerETA,
		Expiry:   expiry,
		Fees:     fees,
	}
}

func (msg *MsgCreateContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

func (msg *MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Consumer)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid consumer address (%s)", err)
	}
	return nil
}
