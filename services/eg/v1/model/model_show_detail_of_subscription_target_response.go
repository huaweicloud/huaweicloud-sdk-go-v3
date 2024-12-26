package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfSubscriptionTargetResponse Response Object
type ShowDetailOfSubscriptionTargetResponse struct {

	// 订阅目标ID
	Id *string `json:"id,omitempty"`

	// 订阅的事件目标名称
	Name *string `json:"name,omitempty"`

	// 订阅的事件目标的提供方类型
	ProviderType *string `json:"provider_type,omitempty"`

	// 订阅的事件目标使用的目标链接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 订阅的事件目标参数列表
	Detail *interface{} `json:"detail,omitempty"`

	KafkaDetail *KafkaTargetDetail `json:"kafka_detail,omitempty"`

	SmnDetail *SmnTargetDetail `json:"smn_detail,omitempty"`

	EgDetail *EgTargetDetail `json:"eg_detail,omitempty"`

	ApigwDetail *ApigwTargetDetail `json:"apigw_detail,omitempty"`

	// 重试次数
	RetryTimes *int32 `json:"retry_times,omitempty"`

	Transform *TransForm `json:"transform,omitempty"`

	DeadLetterQueue *DeadLetterQueue `json:"dead_letter_queue,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfSubscriptionTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfSubscriptionTargetResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfSubscriptionTargetResponse", string(data)}, " ")
}
