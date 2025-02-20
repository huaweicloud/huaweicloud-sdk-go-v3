package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseItemInfo struct {

	// 订阅用户ID。
	Id string `json:"id"`

	// 租户账号ID。
	DomainId string `json:"domain_id"`

	// 订阅用户名称。
	Name string `json:"name"`

	// 订阅用户状态。 UNCONFIRMED：未确认 CONFIRMED：已确认 CANCELLED：已取消
	Status string `json:"status"`

	// 订阅用户分组。
	Group []string `json:"group"`

	// 创建时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime string `json:"create_time"`

	// 更新时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime string `json:"update_time"`

	Http *ListSubscriptionUserResponseHttpEndpointInfo `json:"http,omitempty"`

	Https *ListSubscriptionUserResponseHttpsEndpointInfo `json:"https,omitempty"`

	Sms *ListSubscriptionUserResponseSmsEndpointInfo `json:"sms,omitempty"`

	Email *ListSubscriptionUserResponseEmailEndpointInfo `json:"email,omitempty"`
}

func (o ListSubscriptionUserResponseItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseItemInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseItemInfo", string(data)}, " ")
}
