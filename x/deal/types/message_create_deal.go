package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateDeal = "create_deal"

var _ sdk.Msg = &MsgCreateDeal{}

func NewMsgCreateDeal(creator string, vendor string, commission uint64) *MsgCreateDeal {
	return &MsgCreateDeal{
		Creator:    creator,
		Vendor:     vendor,
		Commission: commission,
	}
}

func (msg *MsgCreateDeal) Route() string {
	return RouterKey
}

func (msg *MsgCreateDeal) Type() string {
	return TypeMsgCreateDeal
}

func (msg *MsgCreateDeal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDeal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDeal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Vendor)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid vendor address (%s)", err)
	}
	return nil
}
