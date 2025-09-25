package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckVulInfo 漏洞信息
type SecurityCheckVulInfo struct {

	// **参数解释**： 漏洞名称 **取值范围**： 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 漏洞ID **取值范围**： 不涉及
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**： 修复紧急度 **取值范围**： 不涉及
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**： 最近扫描时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**： 漏洞类型 **取值范围**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 修复命令行 **取值范围**： 不涉及
	RepairCmd *string `json:"repair_cmd,omitempty"`

	// **参数解释**: 漏洞风险级别 **取值范围**: - Critical：漏洞cvss评分大于等于9；对应控制台页面的高危 - High：漏洞cvss评分大于等于7，小于9；对应控制台页面的中危 - Medium：漏洞cvss评分大于等于4，小于7；对应控制台页面的中危 - Low：漏洞cvss评分小于4；对应控制台页面的低危
	SeverityLevel *string `json:"severity_level,omitempty"`
}

func (o SecurityCheckVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckVulInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckVulInfo", string(data)}, " ")
}
