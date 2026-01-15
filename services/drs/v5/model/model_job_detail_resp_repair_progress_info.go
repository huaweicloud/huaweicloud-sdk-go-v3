package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDetailRespRepairProgressInfo 修复进度明细。
type JobDetailRespRepairProgressInfo struct {

	// 修复状态。 取值：\"FAILED\", \"SUCCEEDED\", \"FINISHED\", \"SUCCESS\"
	Status *string `json:"status,omitempty"`

	// 修复进度，百分比。
	Progress *string `json:"progress,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 总数。
	Count *int32 `json:"count,omitempty"`

	RepairProgressDetails *JobDetailRespRepairProgressInfoRepairProgressDetails `json:"repair_progress_details,omitempty"`
}

func (o JobDetailRespRepairProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailRespRepairProgressInfo struct{}"
	}

	return strings.Join([]string{"JobDetailRespRepairProgressInfo", string(data)}, " ")
}
