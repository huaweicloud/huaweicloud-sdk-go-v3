package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntivirusStatisticResponse Response Object
type ShowAntivirusStatisticResponse struct {

	// **参数解释**： 病毒总数 **取值范围**： 最小值0，最大值2147483647，单位：个
	TotalMalwareNum *int32 `json:"total_malware_num,omitempty"`

	// **参数解释**: 影响主机数量 **取值范围**: 最小值0，最大值2147483647
	MalwareHostNum *int32 `json:"malware_host_num,omitempty"`

	// **参数解释**： 累计扫描任务数 **取值范围**： 最小值0，最大值2147483647，单位：个
	TotalTaskNum *int32 `json:"total_task_num,omitempty"`

	// **参数解释**： 运行中任务数 **取值范围**： 最小值0，最大值2147483647，单位：个
	ScanningTaskNum *int32 `json:"scanning_task_num,omitempty"`

	// **参数解释**： 启动时间 **取值范围**： 最小值0，最大值9223372036854775807；时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算）；单位：ms
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**： 任务类型 **取值范围**： 包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
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
