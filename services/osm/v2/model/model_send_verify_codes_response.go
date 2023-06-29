package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendVerifyCodesResponse Response Object
type SendVerifyCodesResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendVerifyCodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerifyCodesResponse struct{}"
	}

	return strings.Join([]string{"SendVerifyCodesResponse", string(data)}, " ")
}
