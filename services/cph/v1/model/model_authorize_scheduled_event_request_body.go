package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeScheduledEventRequestBody 授权计划事件请求体。
type AuthorizeScheduledEventRequestBody struct {

	// 授权类型。取值范围： maintenance：维护、 redeploy：重部署
	AuthorizationType string `json:"authorization_type"`

	// 计划执行开始时间。仅maintenance类型的事件支持，预约的时间需要比当前时间多5分钟以上，传空字符串表示立即执行。
	NotBefore *string `json:"not_before,omitempty"`
}

func (o AuthorizeScheduledEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeScheduledEventRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeScheduledEventRequestBody", string(data)}, " ")
}
