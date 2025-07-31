package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFileRiskPathsRequest Request Object
type ListIacFileRiskPathsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 文件ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64  **默认取值**: 不涉及
	FileId string `json:"file_id"`

	// **参数解释**: 风险检测规则ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64  **默认取值**: 不涉及
	RuleId string `json:"rule_id"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属的命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o ListIacFileRiskPathsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFileRiskPathsRequest struct{}"
	}

	return strings.Join([]string{"ListIacFileRiskPathsRequest", string(data)}, " ")
}
