package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobSummary 作业数据源。
type JobSummary struct {

	// **参数解释**：训练作业id。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o JobSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSummary struct{}"
	}

	return strings.Join([]string{"JobSummary", string(data)}, " ")
}
