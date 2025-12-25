package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShipperAuthorizationsResponse Response Object
type ListShipperAuthorizationsResponse struct {

	// 错误码，0表示成功，其他值表示错误
	Code *int32 `json:"code,omitempty"`

	Data *ShowShipperAuthorizeResponseBodyData `json:"data,omitempty"`

	// 返回的状态信息，用于描述当前请求的结果
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListShipperAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShipperAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListShipperAuthorizationsResponse", string(data)}, " ")
}
