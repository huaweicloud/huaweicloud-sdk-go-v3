package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeVulScanPolicyRequestInfo struct {

	// **参数解释**: 扫描周期 **约束限制**: 不涉及 **取值范围**: - one_day：每天 - three_day：每三天 - one_week：每周 - one_month：每月 **默认取值**: 不涉及
	ScanPeriod string `json:"scan_period"`

	// **参数解释**: 扫描主机的范围 **约束限制**: 不涉及 **取值范围**: - all_host：扫描全部主机 - specific_host：扫描指定主机 **默认取值**: 不涉及
	ScanRangeType string `json:"scan_range_type"`

	// **参数解释**: 主机ID列表； **约束限制**: 当scan_range_type的值为specific_host时 表示扫描的主机列表 必填 **取值范围**: 最小值0，最大值20000 **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**: \"扫描的漏洞类型列表\" **约束限制**: 不涉及 **取值范围**: 最小值0，最大值5 **默认取值**: 不涉及
	ScanVulTypes *[]string `json:"scan_vul_types,omitempty"`

	// **参数解释**: 扫描策略状态 **约束限制**: 不涉及 **取值范围**: - open : 开启 - close : 关闭  **默认取值** : 不涉及
	Status string `json:"status"`

	Time *ChangeVulScanPolicyRequestInfoTime `json:"time,omitempty"`
}

func (o ChangeVulScanPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulScanPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeVulScanPolicyRequestInfo", string(data)}, " ")
}
