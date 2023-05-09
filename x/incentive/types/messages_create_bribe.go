package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendCreateBribe = "send_create_bribe"

var _ sdk.Msg = &MsgSendCreateBribe{}

func NewMsgSendCreateBribe(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	proposer string,
	title string,
	block uint64,
	denom string,
	amount uint64,
	chain uint64,
) *MsgSendCreateBribe {
	return &MsgSendCreateBribe{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Proposer:         proposer,
		Title:            title,
		Block:            block,
		Denom:            denom,
		Amount:           amount,
		Chain:            chain,
	}
}

func (msg *MsgSendCreateBribe) Route() string {
	return RouterKey
}

func (msg *MsgSendCreateBribe) Type() string {
	return TypeMsgSendCreateBribe
}

func (msg *MsgSendCreateBribe) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendCreateBribe) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendCreateBribe) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}

var (
	_ sdk.Msg = &MsgDistributeBribeRequest{}
)

func (msg *MsgDistributeBribeRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDistributeBribeRequest) Route() string {
	return RouterKey
}

func (msg *MsgDistributeBribeRequest) Type() string {
	return "DistributeBribeRequest"
}

func (msg *MsgDistributeBribeRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgDistributeBribeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func NewMsgDistributeBribeRequest(
	to string,
) *MsgDistributeBribeRequest {
	return &MsgDistributeBribeRequest{
		To: to,
	}
}
