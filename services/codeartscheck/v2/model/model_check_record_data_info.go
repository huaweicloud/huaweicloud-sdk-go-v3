package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckRecordDataInfo struct {

	// 检查任务执行开始时间
	CheckTime *string `json:"check_time,omitempty"`

	// 检查任务执行结束时间
	CheckEndTime *string `json:"check_end_time,omitempty"`

	IssueCounts *CheckRecordIssueCountsInfo `json:"issue_counts,omitempty"`
}

func (o CheckRecordDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordDataInfo struct{}"
	}

	return strings.Join([]string{"CheckRecordDataInfo", string(data)}, " ")
}
