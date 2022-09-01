package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSubscriptionSourceResponse struct {

	// 订阅源ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 订阅的事件源名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 订阅的事件源的提供方类型
	ProviderType *string `json:"provider_type,omitempty" xml:"provider_type"`

	// 订阅的事件源参数列表
	Detail *interface{} `json:"detail,omitempty" xml:"detail"`

	// 订阅事件源的匹配过滤规则
	Filter *interface{} `json:"filter,omitempty" xml:"filter"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 更新时间
	UpdatedTime    *string `json:"updated_time,omitempty" xml:"updated_time"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSubscriptionSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionSourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionSourceResponse", string(data)}, " ")
}
