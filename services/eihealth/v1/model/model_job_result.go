package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResult 作业运行结果信息
type JobResult struct {

	// 输入总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 失败个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 子任务运行时长（秒）。
	SubTasksDuration *[]float32 `json:"sub_tasks_duration,omitempty"`

	// **参数解释**： 分子聚类任务中的分子总数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NumMolecules *int32 `json:"num_molecules,omitempty"`

	// **参数解释**： 聚类成功的分子数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SuccessCount *int32 `json:"success_count,omitempty"`
}

func (o JobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResult struct{}"
	}

	return strings.Join([]string{"JobResult", string(data)}, " ")
}
