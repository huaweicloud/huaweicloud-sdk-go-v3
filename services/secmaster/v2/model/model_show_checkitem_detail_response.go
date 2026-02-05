package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckitemDetailResponse Response Object
type ShowCheckitemDetailResponse struct {

	// 检查项的id
	Uuid *string `json:"uuid,omitempty"`

	// 检查项结果聚合状态
	AggregationHandleStatus *string `json:"aggregation_handle_status,omitempty"`

	// 检查项的检查过程
	AuditProcedure *string `json:"audit_procedure,omitempty"`

	// 检查项的影响
	Impact *string `json:"impact,omitempty"`

	// 检查项所属云服务
	CloudServer *string `json:"cloud_server,omitempty"`

	// 对检查项的描述
	Description *string `json:"description,omitempty"`

	// 表示该检查项的严重程度 informational：提示 low: 低危 medium：中危 high: 高危 fatal：致命
	Level *string `json:"level,omitempty"`

	// 表示该检查项的检查方式 0：手动 1：自动 3: 自动-剧本流程 4: 自动-企业主机安全 5：自动-配置审计服务
	Method *int32 `json:"method,omitempty"`

	// 检查项的名称
	Name *string `json:"name,omitempty"`

	// 表示该检查项的来源 0：自动 2: 自动-剧本流程 3: 手动 4: 自动-企业主机安全 5：自动-配置审计服务
	Source *int32 `json:"source,omitempty"`

	// **参数解释**: 流程ID **约束限制**: 不涉及
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 检查项所属遵从包的信息
	SpecCheckitemList *[]SpecCheckitemModel `json:"spec_checkitem_list,omitempty"`
	HttpStatusCode    int                   `json:"-"`
}

func (o ShowCheckitemDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckitemDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckitemDetailResponse", string(data)}, " ")
}
