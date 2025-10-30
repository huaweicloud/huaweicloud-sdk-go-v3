package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUrgentVulnerabilitiesRequest Request Object
type ListUrgentVulnerabilitiesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 危险程度 **约束限制**: 不涉及 **取值范围**: - Critical：漏洞cvss评分大于等于9；对应控制台页面的高危 - High：漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium：漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low：漏洞cvss评分小于4；对应控制台页面的低危  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：未处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**： 漏洞名称 **取值范围**： 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 是否是容器场景 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： 不涉及
	IsContainer *bool `json:"is_container,omitempty"`

	// **参数解释**： 集群Id **取值范围**： 字符长度0-256位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 扫描任务开始时间的最小值 **取值范围**： 字符长度0-9223372036854775807位
	MinScanTime *int64 `json:"min_scan_time,omitempty"`

	// **参数解释**： 扫描任务开始时间的最大值 **取值范围**： 字符长度0-9223372036854775807位
	MaxScanTime *int64 `json:"max_scan_time,omitempty"`
}

func (o ListUrgentVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUrgentVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListUrgentVulnerabilitiesRequest", string(data)}, " ")
}
