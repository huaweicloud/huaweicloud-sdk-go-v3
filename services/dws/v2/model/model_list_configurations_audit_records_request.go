package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsAuditRecordsRequest Request Object
type ListConfigurationsAuditRecordsRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 任务时间
	ActionTime *int64 `json:"action_time,omitempty"`

	// 过滤配置信息
	FilterBy *string `json:"filter_by,omitempty"`

	// 过滤条件
	Filter *string `json:"filter,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListConfigurationsAuditRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsAuditRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsAuditRecordsRequest", string(data)}, " ")
}
