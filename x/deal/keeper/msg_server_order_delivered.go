package keeper

import (
	"context"
	"example/x/deal/types"
	"strconv"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) OrderDelivered(goCtx context.Context, msg *types.MsgOrderDelivered) (*types.MsgOrderDeliveredResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound

	}

	// validate before processing the message
	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId, msg.ContractId)

	if !found {
		return nil, types.ErrContractNotFound
	}

	if msg.Creator != contract.Consumer {
		return nil, types.ErrInvalidConsumer
	}

	if contract.Status != types.INDELIVERY || contract.Status == types.COMPLETED {
		return nil, types.ErrNotShipped
	}

	startTime, err := time.Parse(types.TIME_FORMAT, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}

	deliveryExpectedTime := startTime.Add(time.Duration(contract.OwnerETA))
	deliveryActualTime := ctx.BlockTime()

	logger.Debug("deliveryExpectedTime :: ", deliveryExpectedTime)
	logger.Debug("deliveryActualTime :: ", deliveryActualTime)

	// calculate the delivery delay in case of any
	if deliveryActualTime.After(deliveryExpectedTime) {
		deliveryTimeDelay := deliveryActualTime.Sub(deliveryExpectedTime).Minutes()
		logger.Debug("deliveryTimeDelay :: ", deliveryTimeDelay)

		if contract.ShippingDelay != 0 {
			// subtract the shipping delay if any
			deliveryTimeDelay = deliveryTimeDelay - float64(contract.ShippingDelay)
			logger.Debug("deliveryTimeDelay after subtracting shipping delay", deliveryTimeDelay)
		}
		contract.DeliveryDelay = uint32(deliveryTimeDelay)
	}

	timeTaken := deliveryActualTime.Sub(startTime).Minutes()
	logger.Debug("timeTaken :: ", timeTaken)

	var refundAmount float64 = 0

	orderPayment := float64(contract.Fees)
	// calculate the penalty for late delivery/shipping for owner & vendor respectively

	if timeTaken != 0 {
		vendorSlashPercent := float64(contract.ShippingDelay) / timeTaken
		logger.Debug("vendorSlashPercent :: ", vendorSlashPercent)

		ownerSlashPercent := float64(contract.DeliveryDelay) / timeTaken
		logger.Debug("ownerSlashPercent :: ", ownerSlashPercent)

		refundAmount = (vendorSlashPercent + ownerSlashPercent) * orderPayment * 0.01
		logger.Debug("refundAmount :: ", refundAmount)
	}

	//deduct delay charges from payment
	totalPay := uint64(orderPayment - refundAmount)
	logger.Debug("TotalPay :: ", totalPay)

	moduleAccount := k.auth.GetModuleAddress(types.ModuleName)
	moduleBalance := k.bank.GetBalance(ctx, moduleAccount, types.TOKEN)
	if moduleBalance.IsLT(contract.GetCoin(totalPay)) {
		panic("Escrow account insufficient balance")
	}

	// calculate the vendor payment for given order based on deal commission
	vendorPay := uint64(0.01 * float64(deal.Commission*totalPay))
	logger.Debug("vendorPay :: ", vendorPay)

	// calculate the owner payment
	ownerPay := totalPay - vendorPay
	logger.Debug("ownerPay :: ", ownerPay)

	// validate the addresses for different parties involved in the contract
	consumerAddress, err := contract.GetConsumerAddress()
	if err != nil {
		panic("Invalid consumer address")
	}

	ownerAddress, err := deal.GetOwnerAddress()
	if err != nil {
		panic("Invalid owner address")
	}

	vendorAddress, err := deal.GetVendorAddress()
	if err != nil {
		panic("Invalid vendor address")
	}

	// process the payment for all the parties
	// refund the delay charges to consumer
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, consumerAddress, sdk.NewCoins(contract.GetCoin(uint64(refundAmount))))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, sdk.NewCoins(contract.GetCoin(ownerPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, vendorAddress, sdk.NewCoins(contract.GetCoin(vendorPay)))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}

	// mark the contract status as completed as order has been delivered to consumer and settlement is complete
	contract.Status = types.COMPLETED
	k.Keeper.SetNewContract(ctx, contract)

	ctx.GasMeter().ConsumeGas(types.SETTLEMENT_GAS, "Order delivered")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.COMPLETED),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
			sdk.NewAttribute(types.CONSUMER, contract.Consumer),
			sdk.NewAttribute(types.OWNER, deal.Owner),
			sdk.NewAttribute(types.VENDOR, deal.Vendor),
			sdk.NewAttribute(types.REFUND_PAY, strconv.FormatUint(uint64(refundAmount), 10)),
			sdk.NewAttribute(types.OWNER_PAY, strconv.FormatUint(ownerPay, 10)),
			sdk.NewAttribute(types.VENDOR_PAY, strconv.FormatUint(vendorPay, 10)),
		),
	)

	return &types.MsgOrderDeliveredResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil

}
