package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentSimpleTicketsResponse Response Object
type ListIncidentSimpleTicketsResponse struct {

	// 服务标识
	ProviderCode string `json:"provider_code"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode string `json:"error_code"`

	// 请求响应描述
	ErrorMsg string `json:"error_msg"`

	// 根据条件查到的事件单信息
	Data           *interface{} `json:"data"`
	HttpStatusCode int          `json:"-"`
}

func (o ListIncidentSimpleTicketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentSimpleTicketsResponse struct{}"
	}

	return strings.Join([]string{"ListIncidentSimpleTicketsResponse", string(data)}, " ")
}
