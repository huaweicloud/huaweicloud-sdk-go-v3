package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRuleResourcesRequest Request Object
type ListCheckRuleResourcesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 类型 **约束限制**: 不涉及 **取值范围**: - image : 镜像  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - cicd：cicd镜像 - registry：仓库镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// **参数解释**: 检查项ID **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	CheckRuleId string `json:"check_rule_id"`

	// **参数解释**: 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	CheckName string `json:"check_name"`

	// **参数解释**: 标准类型 **约束限制**: 不涉及 **取值范围**: - cn_standard：等保合规标准 - hw_standard：云安全实践标准 **默认取值**: 不涉及
	Standard string `json:"standard"`

	// **参数解释**: 检测结果 **约束限制**: 不涉及 **取值范围**: - pass：通过 - failed：未通过  **默认取值**: 不涉及
	ScanResult *string `json:"scan_result,omitempty"`
}

func (o ListCheckRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListCheckRuleResourcesRequest", string(data)}, " ")
}
