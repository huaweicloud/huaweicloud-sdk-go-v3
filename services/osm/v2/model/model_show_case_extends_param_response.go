package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCaseExtendsParamResponse Response Object
type ShowCaseExtendsParamResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCaseExtendsParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCaseExtendsParamResponse struct{}"
	}

	return strings.Join([]string{"ShowCaseExtendsParamResponse", string(data)}, " ")
}
