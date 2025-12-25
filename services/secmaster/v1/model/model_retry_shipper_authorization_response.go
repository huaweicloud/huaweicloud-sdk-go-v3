package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryShipperAuthorizationResponse Response Object
type RetryShipperAuthorizationResponse struct {

	// 错误码，0表示成功，其他值表示错误
	Code *int32 `json:"code,omitempty"`

	// 返回的状态信息，用于描述当前请求的结果
	Msg *string `json:"msg,omitempty"`

	// 投递数据
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o RetryShipperAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryShipperAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"RetryShipperAuthorizationResponse", string(data)}, " ")
}
