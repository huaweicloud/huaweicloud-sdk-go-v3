package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVulScanTaskEstimatedTimeRequestInfo struct {

	// **参数解释**: 漏洞手动检测类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞 - urgent_vul：紧急漏洞  **默认取值**: 不涉及
	ManualScanType *[]string `json:"manual_scan_type,omitempty"`

	// **参数解释**: 检测的主机范围 **约束限制**: 不涉及 **取值范围**: - all_host：全部主机 - specific_host：部分主机  **默认取值**: 不涉及
	RangeType *string `json:"range_type,omitempty"`

	// **参数解释**: 待检测的agent列表 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	AgentIds *[]string `json:"agent_ids,omitempty"`
}

func (o ShowVulScanTaskEstimatedTimeRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulScanTaskEstimatedTimeRequestInfo struct{}"
	}

	return strings.Join([]string{"ShowVulScanTaskEstimatedTimeRequestInfo", string(data)}, " ")
}
