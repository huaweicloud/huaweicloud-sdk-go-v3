package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVariableGroupsRespPipelineVariableGroups struct {

	// 参数组ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 参数组名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// **参数解释**： 参数列表。 **取值范围**： 不涉及。
	Variables *[]QueryVariableGroupDetailRespVariables `json:"variables,omitempty"`

	// 关联的流水线
	RelatedPipelines *[]ListVariableGroupsRespRelatedPipelines `json:"related_pipelines,omitempty"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 更新人ID
	UpdaterId *string `json:"updater_id,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新人名称
	UpdaterName *string `json:"updater_name,omitempty"`

	// 创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int32 `json:"update_time,omitempty"`
}

func (o ListVariableGroupsRespPipelineVariableGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariableGroupsRespPipelineVariableGroups struct{}"
	}

	return strings.Join([]string{"ListVariableGroupsRespPipelineVariableGroups", string(data)}, " ")
}
