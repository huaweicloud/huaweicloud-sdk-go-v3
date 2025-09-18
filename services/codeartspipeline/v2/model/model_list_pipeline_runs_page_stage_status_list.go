package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelineRunsPageStageStatusList struct {

	// **参数解释**： 阶段名称。 **取值范围**： 仅支持输入中文、大小写英文字母、数字、'-'、'_'、','、';'、':'、'.'、'/'、'('、')'、'（'、'）'及空格，其中空格不可在名称开头或结尾使用，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 阶段序列号，0代表第一个流水线阶段。 **取值范围**： 大于等于零。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 流水线阶段状态。 **取值范围**： - INIT：初始化。 - QUEUED：排队。 - RUNNING：运行中。 - CANCELED：取消。 - COMPLETED：已完成。 - FAILED：失败。 - SKIPPED：跳过。 - IGNORED：忽略。 - PAUSED：暂停。 - SUSPEND：挂起。 - ASYNC_RUNNING：异步运行。 - ASYNC_FAILED：异步失败。 - UNSELECTED：未选择。 - REDISPATCH：重新调度。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 阶段开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 阶段结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  阶段ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o ListPipelineRunsPageStageStatusList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsPageStageStatusList struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsPageStageStatusList", string(data)}, " ")
}
