package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriberInstancesRequest Request Object
type ListSubscriberInstancesRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 订阅实例ID
	SubscriberInstanceId *string `json:"subscriber_instance_id,omitempty"`

	// 订阅实例名
	SubscriberInstanceName *string `json:"subscriber_instance_name,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriberInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriberInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriberInstancesRequest", string(data)}, " ")
}
