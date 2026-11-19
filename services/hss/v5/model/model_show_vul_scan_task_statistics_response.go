package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulScanTaskStatisticsResponse Response Object
type ShowVulScanTaskStatisticsResponse struct {

	// **参数解释**: 漏洞扫描任务总次数 **取值范围**: 最小值0，最大值2147483647
	ScanTaskCount *int32 `json:"scan_task_count,omitempty"`

	// **参数解释**: 进行中的漏洞扫描任务次数 **取值范围**: 最小值0，最大值2147483647
	RunningScanTaskCount *int32 `json:"running_scan_task_count,omitempty"`

	LastScanTask   *VulScanTaskStatisticsResponseInfoLastScanTask `json:"last_scan_task,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ShowVulScanTaskStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulScanTaskStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulScanTaskStatisticsResponse", string(data)}, " ")
}
