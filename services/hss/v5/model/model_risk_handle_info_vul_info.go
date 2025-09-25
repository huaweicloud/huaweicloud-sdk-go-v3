package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHandleInfoVulInfo 漏洞信息
type RiskHandleInfoVulInfo struct {

	// **参数解释**: 漏洞扫描任务次数 **取值范围**: 最小值0，最大值2147483647
	VulScanTaskNum *int32 `json:"vul_scan_task_num,omitempty"`

	// **参数解释**: 发现的漏洞数 **取值范围**: 最小值0，最大值2147483647
	VulNum *int64 `json:"vul_num,omitempty"`

	// **参数解释**: 已经处置的漏洞数 **取值范围**: 最小值0，最大值2147483647
	HandledVulNum *int64 `json:"handled_vul_num,omitempty"`

	// **参数解释**: 处置率 **取值范围**: 最小值0，最大值1
	HandledRate *float32 `json:"handled_rate,omitempty"`

	// **参数解释**: 处置率击败的用户率 **取值范围**: 最小值0，最大值1
	BeatRate *float32 `json:"beat_rate,omitempty"`

	// **参数解释**: 新增应急漏洞数 **取值范围**: 最小值0，最大值2147483647
	NewUrgentVulNum *int32 `json:"new_urgent_vul_num,omitempty"`
}

func (o RiskHandleInfoVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHandleInfoVulInfo struct{}"
	}

	return strings.Join([]string{"RiskHandleInfoVulInfo", string(data)}, " ")
}
