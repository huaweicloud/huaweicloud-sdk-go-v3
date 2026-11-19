package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstalledInstanceAgentListRequest Request Object
type InstalledInstanceAgentListRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 必填 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**： 接入后展示维度 **约束限制**： 不涉及 **取值范围**： - workload：工作负载。 - cluster：集群。  **默认取值**： cluster
	Type string `json:"type"`

	// **参数解释**： 工作负载ID，type为cluster时传参为集群ID，展示该集群下工作负载详情；type为workload时传参为工作负载ID，展示该工作负载详情；不传参时，展示type类型所有详情。 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	WorkloadId *string `json:"workload_id,omitempty"`
}

func (o InstalledInstanceAgentListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstalledInstanceAgentListRequest struct{}"
	}

	return strings.Join([]string{"InstalledInstanceAgentListRequest", string(data)}, " ")
}
