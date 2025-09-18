package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelineRunsPagePipelineRuns struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 流水线运行实例ID，[启动流水线](RunPipeline.xml)接口的返回值即为流水线运行实例ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineRunId *string `json:"pipeline_run_id,omitempty"`

	// **参数解释**： 执行人ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	ExecutorId *string `json:"executor_id,omitempty"`

	// **参数解释**： 执行人名称。 **取值范围**： 不涉及。
	ExecutorName *string `json:"executor_name,omitempty"`

	// **参数解释**： 阶段信息列表。 **取值范围**： 不涉及。
	StageStatusList *[]ListPipelineRunsPageStageStatusList `json:"stage_status_list,omitempty"`

	// **参数解释**： 流水线运行实例状态。 **取值范围**： - COMPLETED：已完成。 - RUNNING：运行中。 - FAILED：失败。 - CANCELED：取消。 - PAUSED：暂停。 - SUSPEND：挂起。 - IGNORED：忽略。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 流水线运行序号。 **取值范围**： 大于等于 1。
	RunNumber *int32 `json:"run_number,omitempty"`

	// **参数解释**： 触发类型 **取值范围**： - Manual：手动触发。 - Scheduler：定时任务。 - MR：MR触发。 - Push：Push事件触发。 - CreateTag：Tag事件触发。 - Issue：Issue触发。 - Note：评论触发。
	TriggerType *string `json:"trigger_type,omitempty"`

	BuildParams *ListPipelineRunsPageBuildParams `json:"build_params,omitempty"`

	ArtifactParams *PipelineLatestRunArtifactParams `json:"artifact_params,omitempty"`

	// **参数解释**： 流水线开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 流水线结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 详情页地址，如果x-auth-token中的region信息为空，则详情页地址也返回空。 **取值范围**： 不涉及。
	DetailUrl *string `json:"detail_url,omitempty"`

	// **参数解释**： 修改页地址，如果x-auth-token中的region信息为空，则修改页地址也返回空。 **取值范围**： 不涉及。
	ModifyUrl *string `json:"modify_url,omitempty"`
}

func (o ListPipelineRunsPagePipelineRuns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsPagePipelineRuns struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsPagePipelineRuns", string(data)}, " ")
}
