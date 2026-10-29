package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckRecordDataInfo struct {

	// **参数解释：** 检查任务执行开始时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	CheckTime *string `json:"check_time,omitempty"`

	// **参数解释：** 检查任务执行结束时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
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
