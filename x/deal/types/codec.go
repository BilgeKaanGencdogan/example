package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDeal{}, "deal/CreateDeal", nil)
	cdc.RegisterConcrete(&MsgCreateContract{}, "deal/CreateContract", nil)
	cdc.RegisterConcrete(&MsgCommitContract{}, "deal/CommitContract", nil)
	cdc.RegisterConcrete(&MsgApproveContract{}, "deal/ApproveContract", nil)
	cdc.RegisterConcrete(&MsgShipOrder{}, "deal/ShipOrder", nil)
	cdc.RegisterConcrete(&MsgOrderDelivered{}, "deal/OrderDelivered", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "deal/CancelOrder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShipOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderDelivered{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelOrder{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
