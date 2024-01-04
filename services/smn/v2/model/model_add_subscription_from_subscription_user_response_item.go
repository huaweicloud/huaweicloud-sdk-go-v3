package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSubscriptionFromSubscriptionUserResponseItem struct {

	// 返回码。
	HttpCode int32 `json:"http_code"`

	// 请求的唯一标识ID。
	RequestId string `json:"request_id"`

	// 订阅者的唯一资源标识。
	SubscriptionUrn *string `json:"subscription_urn,omitempty"`

	// 返回信息对应的代码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 服务异常错误信息描述。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AddSubscriptionFromSubscriptionUserResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionFromSubscriptionUserResponseItem struct{}"
	}

	return strings.Join([]string{"AddSubscriptionFromSubscriptionUserResponseItem", string(data)}, " ")
}
