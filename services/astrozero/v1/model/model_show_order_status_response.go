package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderStatusResponse Response Object
type ShowOrderStatusResponse struct {

	// 响应码
	ResCode *string `json:"resCode,omitempty"`

	// 响应信息
	ResMsg *string `json:"resMsg,omitempty"`

	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowOrderStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowOrderStatusResponse", string(data)}, " ")
}
