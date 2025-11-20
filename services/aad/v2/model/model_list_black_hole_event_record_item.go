package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlackHoleEventRecordItem item
type ListBlackHoleEventRecordItem struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 租户名
	TenantName string `json:"tenant_name"`

	// 高防ip
	Vip string `json:"vip"`

	// ip id
	VipId string `json:"vip_id"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 事件类型 block-黑洞中，unblock-黑洞结束
	EventType string `json:"event_type"`

	// 封堵开始时间
	StartTime int32 `json:"start_time"`

	// 封堵结束时间
	EndTime *int32 `json:"end_time,omitempty"`
}

func (o ListBlackHoleEventRecordItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlackHoleEventRecordItem struct{}"
	}

	return strings.Join([]string{"ListBlackHoleEventRecordItem", string(data)}, " ")
}
