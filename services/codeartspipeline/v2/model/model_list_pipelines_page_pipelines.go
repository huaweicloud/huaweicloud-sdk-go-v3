package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelinesPagePipelines struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 流水线名称。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**： 组件ID。 **取值范围**： 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**： 是否为变更流水线。 **取值范围**： - true：是变更流水线。 - false：不是变更流水线。
	IsPublish *bool `json:"is_publish,omitempty"`

	// **参数解释**： 是否收藏此流水线。 **取值范围**： - true：已收藏流水线。 - false：未收藏流水线。
	IsCollect *bool `json:"is_collect,omitempty"`

	// **参数解释**： 流水线版本。 **取值范围**： 默认3.0。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	LatestRun *ListPipelinesPageLatestRun `json:"latest_run,omitempty"`

	// **参数解释**： 旧版转新版标识。 **取值范围**： 不涉及。
	ConvertSign *int32 `json:"convert_sign,omitempty"`

	// **参数解释**： 流水线涉密等级。 **取值范围**： 正整数。 null：未设置密级。 1：最低密级。
	SecurityLevel *int32 `json:"security_level,omitempty"`
}

func (o ListPipelinesPagePipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPagePipelines struct{}"
	}

	return strings.Join([]string{"ListPipelinesPagePipelines", string(data)}, " ")
}
