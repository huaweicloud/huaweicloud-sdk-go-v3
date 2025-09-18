package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineRunDetailResponse Response Object
type ShowPipelineRunDetailResponse struct {

	// **参数解释**： 流水线运行实例ID，[启动流水线](RunPipeline.xml)接口的返回值即为流水线运行实例ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 流水线版本。 **取值范围**： 默认3.0。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 流水线名称。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 流水线运行描述。 **取值范围**： 最长1024字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否为变更流水线。 **取值范围**： - true：是变更流水线。 - false：不是变更流水线。
	IsPublish *bool `json:"is_publish,omitempty"`

	// **参数解释**： 运行人ID，用户的userId。 **取值范围**： 32位字符，仅由数字和字母组成。
	ExecutorId *string `json:"executor_id,omitempty"`

	// **参数解释**： 运行人名称。 **取值范围**： 不涉及。
	ExecutorName *string `json:"executor_name,omitempty"`

	// **参数解释**： 流水线运行实例状态。 **取值范围**： - COMPLETED：已完成。 - RUNNING：运行中。 - FAILED：失败。 - CANCELED：取消。 - PAUSED：暂停。 - SUSPEND：挂起。 - IGNORED：忽略。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 流水线触发类型。 - Manual：手动触发。 - Scheduler：定时任务。 - MR：MR触发。 - Push：Push事件触发。 - CreateTag：Tag事件触发。 - Issue：Issue触发。 - Note：评论触发。 **取值范围**： 不涉及。
	TriggerType *string `json:"trigger_type,omitempty"`

	// **参数解释**： 流水线运行序号。 **取值范围**： 大于等于 1。
	RunNumber *int32 `json:"run_number,omitempty"`

	// **参数解释**： 流水线开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 流水线结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 阶段运行信息列表，包含各个阶段的详细运行信息。 **约束限制**： 不涉及。
	Stages *[]StageRun `json:"stages,omitempty"`

	// **参数解释**： 租户ID，用户账号的domainId。 **取值范围**： 32位字符，仅由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 局点。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 微服务ID。可以通过[查询微服务列表](ListMicroservice.xml)接口获取，其中data.id即为微服务ID。 **取值范围**： 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**： 语言，暂时仅包含中英文。 **取值范围**： zh-cn, en-us。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 流水线执行源信息。 **取值范围**： 不涉及。
	Sources *[]RunPipelineSource `json:"sources,omitempty"`

	// **参数解释**： 流水线运行产物。 **取值范围**： 不涉及。
	Artifacts *[]PackageInfo `json:"artifacts,omitempty"`

	// **参数解释**： 流水线运行实例ID，[启动流水线](RunPipeline.xml)接口的返回值即为流水线运行实例ID。 **取值范围**： 不涉及。
	SubjectId *string `json:"subject_id,omitempty"`

	// **参数解释**： 分组ID。 **取值范围**： 32位字符，由数字和字母组成。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 分组名称。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 详情页地址。 **取值范围**： 不涉及。
	DetailUrl *string `json:"detail_url,omitempty"`

	// **参数解释**： 当前系统时间。 **取值范围**： 不涉及。
	CurrentSystemTime *int64 `json:"current_system_time,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowPipelineRunDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineRunDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineRunDetailResponse", string(data)}, " ")
}
