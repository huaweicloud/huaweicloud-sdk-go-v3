package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineBuildResult struct {

	// **参数解释**： 流水线执行ID。 **取值范围**： 不涉及。
	BuildId string `json:"build_id"`

	// **参数解释**： 运行耗时，单位为毫秒。 **取值范围**： 不涉及。
	ElapseTime *string `json:"elapse_time,omitempty"`

	// **参数解释**： 执行结束时间。 **取值范围**： 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**： 运行结果。 **取值范围**： - success：成功。 - error：失败。 - aborted：终止。
	Outcome string `json:"outcome"`

	// **参数解释**： 流水线id。 **取值范围**： 32位字符，由数字和字母组成。
	PipelineId string `json:"pipeline_id"`

	// **参数解释**： 流水线名称。 **取值范围**： 不涉及。
	PipelineName string `json:"pipeline_name"`

	// **参数解释**： 执行开始时间。 **取值范围**： 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**： 运行状态。 **取值范围**： - waiting：等待中。 - running：运行中。 - verifying：待审核。 - suspending：挂起。 - completed：执行完成。
	Status string `json:"status"`
}

func (o PipelineBuildResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBuildResult struct{}"
	}

	return strings.Join([]string{"PipelineBuildResult", string(data)}, " ")
}
