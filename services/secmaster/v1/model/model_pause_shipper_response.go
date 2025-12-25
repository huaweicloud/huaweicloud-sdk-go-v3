package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseShipperResponse Response Object
type PauseShipperResponse struct {

	// 错误码，0表示成功，其他值表示错误
	Code *int32 `json:"code,omitempty"`

	// 返回的状态信息，用于描述当前请求的结果
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PauseShipperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseShipperResponse struct{}"
	}

	return strings.Join([]string{"PauseShipperResponse", string(data)}, " ")
}
