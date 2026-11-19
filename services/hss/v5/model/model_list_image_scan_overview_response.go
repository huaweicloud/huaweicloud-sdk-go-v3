package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageScanOverviewResponse Response Object
type ListImageScanOverviewResponse struct {

	// **参数解释**: 镜像总数 **取值范围**: 最小值0，最大值2147483647
	TotalImageNum *int32 `json:"total_image_num,omitempty"`

	// **参数解释**: 存在风险的镜像数 **取值范围**: 最小值0，最大值2147483647
	RiskyImageNum *int32 `json:"risky_image_num,omitempty"`

	// **参数解释**: 漏洞数量 **取值范围**: 最小值0，最大值2147483647
	VulsNum *int32 `json:"vuls_num,omitempty"`

	// **参数解释**: 风险配置数 **取值范围**: 最小值0，最大值2147483647
	RiskyConfigNum *int32 `json:"risky_config_num,omitempty"`

	// **参数解释**: 恶意文件数 **取值范围**: 最小值0，最大值2147483647
	MalwaresNum *int32 `json:"malwares_num,omitempty"`

	// **参数解释**: 未扫描数 **取值范围**: 最小值0，最大值2147483647
	UnscannedNum *int32 `json:"unscanned_num,omitempty"`

	// **参数解释**: 扫描中任务数 **取值范围**: 最小值0，最大值2147483647
	ScanningTaskNum *int32 `json:"scanning_task_num,omitempty"`

	// **参数解释**: 已完成任务数 **取值范围**: 最小值0，最大值2147483647
	FinishedTaskNum *int32 `json:"finished_task_num,omitempty"`

	// **参数解释**: 最近扫描时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 本地扫描中任务数 **取值范围**: 最小值0，最大值2147483647
	LocalScanningTaskNum *int32 `json:"local_scanning_task_num,omitempty"`

	// **参数解释**: 本地已完成任务数 **取值范围**: 最小值0，最大值2147483647
	LocalFinishedTaskNum *int32 `json:"local_finished_task_num,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ListImageScanOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageScanOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListImageScanOverviewResponse", string(data)}, " ")
}
