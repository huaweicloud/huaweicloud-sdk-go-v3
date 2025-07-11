package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTicketActionResponse Response Object
type ExecuteTicketActionResponse struct {

	// 服务标识
	ProviderCode *string `json:"provider_code,omitempty"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode *string `json:"error_code,omitempty"`

	// 请求响应描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 响应数据
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ExecuteTicketActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTicketActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTicketActionResponse", string(data)}, " ")
}
