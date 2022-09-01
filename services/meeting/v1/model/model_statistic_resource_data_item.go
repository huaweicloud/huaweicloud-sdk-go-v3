package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议已购资源使用数据的单个时间点数据。
type StatisticResourceDataItem struct {

	// 日期/月份。
	Time *string `json:"time,omitempty" xml:"time"`

	// VMR方数。 category = used_vmr_info时有效。
	VmrParties *string `json:"vmrParties,omitempty" xml:"vmrParties"`

	// VMR并发使用数。 category = used_vmr_info时有效。
	MaxConcurrencyVmrCount *string `json:"maxConcurrencyVmrCount,omitempty" xml:"maxConcurrencyVmrCount"`

	// 直播端口并发使用数。 category = used_live_info时有效。
	LivePortUsedCount *string `json:"livePortUsedCount,omitempty" xml:"livePortUsedCount"`

	// 录播使用空间(G)。 category = used_record_info时有效。
	RecordUsedSize *string `json:"recordUsedSize,omitempty" xml:"recordUsedSize"`

	// PSTN外呼时长(分钟)。 category = used_pstn_info时有效。
	PstnUsedDuration *string `json:"pstnUsedDuration,omitempty" xml:"pstnUsedDuration"`
}

func (o StatisticResourceDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticResourceDataItem struct{}"
	}

	return strings.Join([]string{"StatisticResourceDataItem", string(data)}, " ")
}
