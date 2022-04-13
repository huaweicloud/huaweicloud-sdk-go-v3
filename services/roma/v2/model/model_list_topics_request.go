package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTopicsRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备ID

	DeviceId int32 `json:"device_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
	// topic名称

	Name *string `json:"name,omitempty"`
	// topic权限，0为发布，1为订阅

	TopicPermission *int32 `json:"topic_permission,omitempty"`
	// topic类型，0为设备类topic，1为产品类topic

	TopicType *int32 `json:"topic_type,omitempty"`
	// topic是否为自定义，0为基础topic，1为自定义topic

	IsPrivate *int32 `json:"is_private,omitempty"`
}

func (o ListTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicsRequest", string(data)}, " ")
}
