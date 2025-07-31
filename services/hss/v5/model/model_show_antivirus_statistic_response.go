package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntivirusStatisticResponse Response Object
type ShowAntivirusStatisticResponse struct {

	// 病毒总数
	TotalMalwareNum *int32 `json:"total_malware_num,omitempty"`

	// **参数解释**: 影响主机数量 **取值范围**: 最小值0，最大值2147483647
	MalwareHostNum *int32 `json:"malware_host_num,omitempty"`

	// 累计扫描任务数
	TotalTaskNum *int32 `json:"total_task_num,omitempty"`

	// 运行中任务数
	ScanningTaskNum *int32 `json:"scanning_task_num,omitempty"`

	// 启动时间，毫秒
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType       *string `json:"scan_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAntivirusStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntivirusStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowAntivirusStatisticResponse", string(data)}, " ")
}
