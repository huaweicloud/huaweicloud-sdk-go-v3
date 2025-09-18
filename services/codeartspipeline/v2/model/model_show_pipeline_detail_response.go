package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineDetailResponse Response Object
type ShowPipelineDetailResponse struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 流水线名称。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 对流水线的补充描述。 **取值范围**： 不超过1024字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 流水线版本，默认为3.0。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 当前环境所属局点。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 所属租户ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 所属微服务ID。可以通过[查询微服务列表](ListMicroservice.xml)接口获取，其中data.id即为微服务ID。 **取值范围**： 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**： 是否为变更流水线。 **取值范围**： - true：是变更流水线。 - false：不是变更流水线。
	IsPublish *bool `json:"is_publish,omitempty"`

	// **参数解释**： 流水线创建人ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 流水线创建人名称。 **取值范围**： 不涉及。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**： 流水线上次更新人ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	UpdaterId *string `json:"updater_id,omitempty"`

	// **参数解释**： 流水线创建时间。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 流水线更新时间。 **取值范围**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 流水线是否被当前用户收藏。 **取值范围**： - true：流水线已被收藏。 - false：流水线未被收藏。
	IsCollect *bool `json:"is_collect,omitempty"`

	// **参数解释**： 流水线源列表。 **取值范围**： 不涉及。
	Sources *[]PipelineSource `json:"sources,omitempty"`

	// **参数解释**： 流水线自定义参数。 **取值范围**： 不涉及。
	Variables *[]PipelineVariable `json:"variables,omitempty"`

	// **参数解释**： 流水线定时任务设置。 **取值范围**： 不涉及。
	Schedules *[]PipelineSchedule `json:"schedules,omitempty"`

	// **参数解释**： 流水线事件触发设置。 **取值范围**： 不涉及。
	Triggers *[]PipelineTrigger `json:"triggers,omitempty"`

	// **参数解释**： 流水线所属分组ID。 **取值范围**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 流水线定义JSON。 **取值范围**： 不涉及。
	Definition *string `json:"definition,omitempty"`

	// **参数解释**： 流水线涉密等级。 **取值范围**： 不涉及。
	SecurityLevel  *int32 `json:"security_level,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPipelineDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineDetailResponse", string(data)}, " ")
}
