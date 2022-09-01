package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseInfoResp struct {

	// 响应的HTTP状态码
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 响应的Body模板
	Body *string `json:"body,omitempty" xml:"body"`

	// 是否为默认响应
	Default *bool `json:"default,omitempty" xml:"default"`
}

func (o ResponseInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseInfoResp struct{}"
	}

	return strings.Join([]string{"ResponseInfoResp", string(data)}, " ")
}
