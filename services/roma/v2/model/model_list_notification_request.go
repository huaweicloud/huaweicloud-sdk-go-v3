package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNotificationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，分页的偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int32 `json:"offset,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`
}

func (o ListNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationRequest", string(data)}, " ")
}
