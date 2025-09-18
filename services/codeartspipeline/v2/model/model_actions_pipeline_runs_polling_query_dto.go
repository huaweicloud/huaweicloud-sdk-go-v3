package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionsPipelineRunsPollingQueryDto struct {

	// **参数解释**： 流水线运行实例ID列表。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	PipelineRunIds *[]string `json:"pipeline_run_ids,omitempty"`
}

func (o ActionsPipelineRunsPollingQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionsPipelineRunsPollingQueryDto struct{}"
	}

	return strings.Join([]string{"ActionsPipelineRunsPollingQueryDto", string(data)}, " ")
}
