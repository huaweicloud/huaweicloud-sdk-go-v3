package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTicketResponse Response Object
type CreateTicketResponse struct {

	// 服务标识
	ProviderCode string `json:"provider_code"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode string `json:"error_code"`

	// 请求响应描述
	ErrorMsg string `json:"error_msg"`

	// 返回体
	Data           *interface{} `json:"data"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateTicketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTicketResponse struct{}"
	}

	return strings.Join([]string{"CreateTicketResponse", string(data)}, " ")
}
