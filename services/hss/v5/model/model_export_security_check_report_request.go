package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSecurityCheckReportRequest Request Object
type ExportSecurityCheckReportRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器组ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** : \"配置检查（基线）的名称，例如SSH、CentOS 7、Windows\"。 **约束限制** : 不涉及 **取值范围** : - cn_standard: 等保合规标准 - hw_standard: 云安全实践标准 **默认取值** : 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**: 标准类型。 **约束限制**: 不涉及 **取值范围**: - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 **默认取值**: 不涉及
	Standard *string `json:"standard,omitempty"`

	// **参数解释**: 检查结果。 **约束限制**: 不涉及 **取值范围**: - pass : 检查通过 - failed : 检查未通过 **默认取值**: 不涉及
	ScanResult *string `json:"scan_result,omitempty"`

	// **参数解释**: 风险级别。 **约束限制**: 不涉及 **取值范围**: - Low : 低危 - Medium : 中危 - High : 高危 **默认取值**: 不涉及
	Severity *string `json:"severity,omitempty"`
}

func (o ExportSecurityCheckReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSecurityCheckReportRequest struct{}"
	}

	return strings.Join([]string{"ExportSecurityCheckReportRequest", string(data)}, " ")
}
