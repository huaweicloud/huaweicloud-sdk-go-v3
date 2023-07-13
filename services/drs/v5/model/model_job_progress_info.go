package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobProgressInfo 任务进度信息。
type JobProgressInfo struct {

	// 迁移对比百分比。
	Progress *string `json:"progress,omitempty"`

	// 增量迁移时延（单位：s）。
	IncrTransDelay *string `json:"incr_trans_delay,omitempty"`

	// 增量迁移时延（单位：ms）。
	IncrTransDelayMillis *string `json:"incr_trans_delay_millis,omitempty"`

	// 迁移模式。
	TaskMode *string `json:"task_mode,omitempty"`

	// 迁移状态。
	TransferStatus *string `json:"transfer_status,omitempty"`

	// 迁移时间。
	ProcessTime *string `json:"process_time,omitempty"`

	// 预计剩余时间。
	RemainingTime *string `json:"remaining_time,omitempty"`

	// 全量迁移进度详情。
	ProgressMap map[string]ProgressCompleteInfo `json:"progress_map,omitempty"`
}

func (o JobProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobProgressInfo struct{}"
	}

	return strings.Join([]string{"JobProgressInfo", string(data)}, " ")
}
