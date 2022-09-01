package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSubscriptionTargetResponse struct {

	// 订阅目标ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 订阅的事件目标名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 订阅的事件目标的提供方类型
	ProviderType *string `json:"provider_type,omitempty" xml:"provider_type"`

	// 订阅的事件目标参数列表
	Detail *interface{} `json:"detail,omitempty" xml:"detail"`

	Transform *SubscriptionTargetInfoTransform `json:"transform,omitempty" xml:"transform"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 更新时间
	UpdatedTime    *string `json:"updated_time,omitempty" xml:"updated_time"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSubscriptionTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTargetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTargetResponse", string(data)}, " ")
}
