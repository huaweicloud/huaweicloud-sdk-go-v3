package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerlessVulStatisticsResponse Response Object
type ShowServerlessVulStatisticsResponse struct {

	// **参数解释**: 应用漏洞总数 **取值范围**: 最小值0，最大值2147483647
	TotalAppVulNum *int32 `json:"total_app_vul_num,omitempty"`

	// **参数解释**: 高风险应用漏洞数 **取值范围**: 最小值0，最大值2147483647
	HighRiskyAppVul *int32 `json:"high_risky_app_vul,omitempty"`

	// **参数解释**: linux漏洞总数 **取值范围**: 最小值0，最大值2147483647
	TotalLinuxVulNum *int32 `json:"total_linux_vul_num,omitempty"`

	// **参数解释**: 高风险应用漏洞数 **取值范围**: 最小值0，最大值2147483647
	HighRiskyLinuxVul *int32 `json:"high_risky_linux_vul,omitempty"`

	// **参数解释**: 存在漏洞的实例数 **取值范围**: 最小值0，最大值2147483647
	RiskyPodNum *int32 `json:"risky_pod_num,omitempty"`

	// **参数解释**: 累计手动漏洞扫描次数 **取值范围**: 最小值0，最大值2147483647
	TotalScanNum *int32 `json:"total_scan_num,omitempty"`

	// **参数解释**: 最近检测时间 **取值范围**: 最小值0，最大值4071095999000
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// **参数解释**: 正在扫描数 **取值范围**: 最小值0，最大值4071095999000
	ScanningTaskNum *int64 `json:"scanning_task_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowServerlessVulStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerlessVulStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowServerlessVulStatisticsResponse", string(data)}, " ")
}
