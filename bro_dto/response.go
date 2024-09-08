package bro_dto

import (
	"fmt"

	"github.com/bot-reserve-order/bro-common-lib/bro_domains"
	"github.com/bot-reserve-order/bro-common-lib/util"
	"github.com/shopspring/decimal"
)

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TxnID   string `json:"txn_id"`
}

type CommonErrorResponse struct {
	OriginalError string `json:"original_error,omitempty"`
	Location      string `json:"location,omitempty"`
}

type BroErrorResponse struct {
	CommonResponse
	*CommonErrorResponse
}

func (be BroErrorResponse) Error() string {
	be.Location = util.Caller(2)
	return fmt.Sprint(be.Code, " : ", be.Message, " : ", be.Location)
}

func NewSuccessCommonResponse() CommonResponse {
	return CommonResponse{
		Code:    0,
		Message: "success",
	}
}

type ResponseLogin struct {
	CommonResponse
	Data string `json:"data"`
}

type ResponseCreateCondition struct {
	CommonResponse
	Condition bro_domains.Conditions `json:"data"`
}

type ResponseGetProfile struct {
	CommonResponse
	Data bro_domains.User `json:"data"`
}

type ResponseGetConditions struct {
	CommonResponse
	Data ResponseGetConditionsData `json:"data"`
}

type ResponseGetConditionsData struct {
	Conds     []bro_domains.Conditions `json:"conds"`
	Page      uint                     `json:"page"`
	PerPage   uint                     `json:"per_page"`
	TotalPage uint                     `json:"total_page"`
}

type ResponseCreateRoomNotify struct {
	CommonResponse
	Data bro_domains.RoomNotify `json:"data"`
}

type ResponseGetCredit struct {
	CommonResponse
	Data decimal.Decimal `json:"data"`
}

type ResponseCondition struct {
	CommonResponse
	Data bro_domains.Conditions `json:"data"`
}

type ResponseSumCredit struct {
	CommonResponse
	Data decimal.Decimal `json:"data"`
}

type ResponseRoomNotify struct {
	CommonResponse
	Data []bro_domains.RoomNotify `json:"data"`
}

type ResponseQrDeposit struct {
	CommonResponse
	Data string `json:"data"`
}

type ResponseGetUsers struct {
	CommonResponse
	Data []bro_domains.User `json:"data"`
}

type ResponseGetRequestWithdraw struct {
	CommonResponse
	Data []bro_domains.WithdrawRequest `json:"data"`
}

type ResponseCreateWithdrawRequest struct {
	CommonResponse
	Data bro_domains.WithdrawRequest `json:"data"`
}

type ResponseGetCreditsData struct {
	Conds     []bro_domains.Credit `json:"credits"`
	Page      uint                 `json:"page"`
	PerPage   uint                 `json:"per_page"`
	TotalPage uint                 `json:"total_page"`
}

type ResponseGetCredits struct {
	CommonResponse
	Data ResponseGetCreditsData `json:"data"`
}
