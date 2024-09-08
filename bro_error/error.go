package bro_error

import (
	"context"

	"github.com/bot-reserve-order/bro-common-lib/bro_dto"
	"github.com/bot-reserve-order/bro-common-lib/util"
)

func newError(ctx context.Context, code int, msg string, origin error) *bro_dto.BroErrorResponse {
	trace, _ := ctx.Value("trace_id").(string)

	err := &bro_dto.BroErrorResponse{
		CommonResponse: bro_dto.CommonResponse{
			Code:    code,
			Message: msg,
			TxnID:   trace,
		},
	}
	if origin != nil {
		err.CommonErrorResponse = &bro_dto.CommonErrorResponse{
			OriginalError: origin.Error(),
			Location:      util.Caller(3),
		}
	}
	return err
}

func NewBroNoError(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 0, "success", nil)
}

func NewBroInternalServerError(ctx context.Context, origin error) *bro_dto.BroErrorResponse {
	return newError(ctx, 500, "Inter server error", origin)
}

func NewBroBadRequestError(ctx context.Context, origin error) *bro_dto.BroErrorResponse {
	return newError(ctx, 400, "Bad request", origin)
}

func NewBroForbidenError(ctx context.Context, origin error) *bro_dto.BroErrorResponse {
	return newError(ctx, 403, "Forbiden request", origin)
}

// 1000
func NewBroErrorUsernameExist(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1000, "Username already exist", nil)
}

func NewBroErrorInvalidReference(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1001, "Invalid referencce", nil)
}

func NewBroErrorCreateUser(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1003, "Register user error", nil)
}

func NewBroErrorInvalidUserNameOrPassword(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1003, "Invalid username or password", nil)
}

func NewBroErrorWalletExist(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1004, "Wallet already exist", nil)
}

func NewBroErrorInvalidOldPassword(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1005, "Invalid old password", nil)
}

func NewBroErrorUpdatePassword(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1006, "Update password error", nil)
}

func NewBroErrorCustomIdAlreadyExist(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1007, "Custom ID already exist", nil)
}

func NewBroErrorUserDoseNotExist(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1008, "User dose not exist", nil)
}

func NewBroErrorUpdateUserStatus(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1009, "Update user status error", nil)
}

func NewBroErrorUpdateUserLineToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1010, "Update line token error", nil)
}

func NewBroErrorUpdateFleetToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1011, "Update fleet token error", nil)
}

func NewBroErrorUpdateUser(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1012, "Update user error", nil)
}

func NewBroErrorRequireLineToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 1013, "Require Line notify token", nil)
}

// 2000
func NewBroErrorGetFleetProfileError(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 2000, "Get fleet profile error", nil)
}

func NewBroErrorGetFleetProfile(ctx context.Context, msg string) *bro_dto.BroErrorResponse {
	return newError(ctx, 2001, msg, nil)
}

func NewBroErrorFleetLoginError(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 2002, "Fleet login error", nil)
}

func NewBroErrorFleetLogin(ctx context.Context, msg string) *bro_dto.BroErrorResponse {
	return newError(ctx, 2003, msg, nil)
}

func NewBroErrorGetStaffInfoError(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 2004, "Get staff info error", nil)
}

func NewBroErrorGetOrderHistoryError(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 2004, "Get order history error", nil)
}

// 3000
func NewBroErrorCreateToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3000, "Create credential error", nil)
}

func NewBroErrorVerifyTokenAndCheckRevokedToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3001, "Verify and check credential error", nil)
}

func NewBroErrorVerifyToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3002, "Verify credential error", nil)
}

func NewBroErrorCreateCookie(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3003, "Create cookie error", nil)
}

func NewBroErrorVerifyCookieAndCheckRevokedToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3004, "Verify and check cookie error", nil)
}

func NewBroErrorVerifyCookie(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 3005, "Verify cookie error", nil)
}

// 4000
func NewBroErrorCreateRoomNotify(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 4000, "Verify and check cookie error", nil)
}

func NewBroErrorGetRoomNotify(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 4001, "Get room notify error", nil)
}

func NewBroErrorRoomNotifyNotFound(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 4002, "Room notify not found", nil)
}

func NewBroErrorDeleteRoomNotify(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 4003, "Delete room notify error", nil)
}

// 5000
func NewBroErrorInvalidCartype(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5000, "Invalid car type", nil)
}
func NewBroErrorCreateCondition(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5001, "Create condition error", nil)
}
func NewBroErrorConditionNotfound(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5002, "Condition not found", nil)
}
func NewBroErrorDeleteCondition(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5003, "Delete condition error", nil)
}
func NewBroErrorGetCondition(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5004, "Get condition error", nil)
}
func NewBroErrorUpdateCondition(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5005, "Update condition error", nil)
}
func NewBroErrorSubmitInvalidCondition(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5006, "Submit invalid condition", nil)
}
func NewBroErrorConditionAlreadySubmit(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5007, "Condition already submit", nil)
}
func NewBroErrorConditionNotMatchOrder(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 5008, "Condition not match with order", nil)
}

// 6000
func NewBroErrorCreateCredit(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 6000, "Create credit error", nil)
}
func NewBroErrorGetSumCredit(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 6001, "Get sum credit error", nil)
}
func NewBroErrorInsufficientCredit(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 6002, "Insufficient credit", nil)
}
func NewBroErrorNegativeCredit(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 6003, "Negative credit", nil)
}
func NewBroErrorInvalidDepositTrueWalletToken(ctx context.Context) *bro_dto.BroErrorResponse {
	return newError(ctx, 6004, "Invalid token", nil)
}
