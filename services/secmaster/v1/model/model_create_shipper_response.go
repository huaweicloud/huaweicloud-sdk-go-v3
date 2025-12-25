package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShipperResponse Response Object
type CreateShipperResponse struct {

	// 错误码，0表示成功，其他值表示错误
	Code *int32 `json:"code,omitempty"`

	// 返回的数据
	Data *int32 `json:"data,omitempty"`

	// 返回的状态信息
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateShipperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperResponse struct{}"
	}

	return strings.Join([]string{"CreateShipperResponse", string(data)}, " ")
}
