package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperResponse Response Object
type ShowShipperResponse struct {

	// 状态码
	Code *int32 `json:"code,omitempty"`

	Data *ShowShipperResponseBodyData `json:"data,omitempty"`

	// 响应消息
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowShipperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperResponse struct{}"
	}

	return strings.Join([]string{"ShowShipperResponse", string(data)}, " ")
}
