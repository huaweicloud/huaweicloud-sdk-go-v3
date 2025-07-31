package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountBaseResponse 基本响应体，需要增加data属性
type AccountBaseResponse struct {

	// 服务标识
	ProviderCode string `json:"provider_code"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode string `json:"error_code"`

	// 请求响应描述
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AccountBaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountBaseResponse struct{}"
	}

	return strings.Join([]string{"AccountBaseResponse", string(data)}, " ")
}
