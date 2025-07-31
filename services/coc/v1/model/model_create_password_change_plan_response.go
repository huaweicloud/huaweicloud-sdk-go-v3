package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePasswordChangePlanResponse Response Object
type CreatePasswordChangePlanResponse struct {

	// 服务标识
	ProviderCode *string `json:"provider_code,omitempty"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode *string `json:"error_code,omitempty"`

	// 请求响应描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *CreatePasswordChangePlanResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o CreatePasswordChangePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordChangePlanResponse struct{}"
	}

	return strings.Join([]string{"CreatePasswordChangePlanResponse", string(data)}, " ")
}
