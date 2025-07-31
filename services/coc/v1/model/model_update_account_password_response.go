package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountPasswordResponse Response Object
type UpdateAccountPasswordResponse struct {

	// 服务标识
	ProviderCode string `json:"provider_code"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode string `json:"error_code"`

	// 请求响应描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 响应数据
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAccountPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccountPasswordResponse", string(data)}, " ")
}
