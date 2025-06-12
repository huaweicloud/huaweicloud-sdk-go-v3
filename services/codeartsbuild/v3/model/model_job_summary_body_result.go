package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobSummaryBodyResult 结果
type JobSummaryBodyResult struct {
	Summary *JobSummary `json:"summary,omitempty"`
}

func (o JobSummaryBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSummaryBodyResult struct{}"
	}

	return strings.Join([]string{"JobSummaryBodyResult", string(data)}, " ")
}
