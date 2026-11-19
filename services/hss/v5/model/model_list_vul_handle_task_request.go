package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHandleTaskRequest Request Object
type ListVulHandleTaskRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 处置任务类型 **约束限制**: 不涉及 **取值范围**: - immediate_repair：立即修复。 - verify：验证。  **默认取值**: 不涉及
	HandleType *string `json:"handle_type,omitempty"`

	// **参数解释**: 任务开始时间的最小值，时间戳，单位毫秒 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	MinStartTime *int64 `json:"min_start_time,omitempty"`

	// **参数解释**: 任务开始时间的最大值，时间戳，单位毫秒 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	MaxStartTime *int64 `json:"max_start_time,omitempty"`

	// **参数解释**: 任务状态 **约束限制**: 不涉及 **取值范围**: - running：执行中。 - finished：已完成。  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`
}

func (o ListVulHandleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHandleTaskRequest struct{}"
	}

	return strings.Join([]string{"ListVulHandleTaskRequest", string(data)}, " ")
}
