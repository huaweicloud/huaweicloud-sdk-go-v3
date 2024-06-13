package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemConfigsResponse Response Object
type ShowSystemConfigsResponse struct {

	// 接口调用失败错误码
	Code *string `json:"code,omitempty"`

	Data *[]SystemConfig `json:"data,omitempty"`

	// 接口调用错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSystemConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowSystemConfigsResponse", string(data)}, " ")
}
