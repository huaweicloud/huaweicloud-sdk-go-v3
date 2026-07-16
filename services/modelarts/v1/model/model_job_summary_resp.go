package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobSummaryResp 作业数据源。
type JobSummaryResp struct {

	// **参数解释**：训练作业id。 **取值范围**：不涉及。
	JobId string `json:"job_id"`
}

func (o JobSummaryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSummaryResp struct{}"
	}

	return strings.Join([]string{"JobSummaryResp", string(data)}, " ")
}
