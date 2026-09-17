package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVariableGroupsResponse Response Object
type ListVariableGroupsResponse struct {

	// 详情列表
	PipelineVariableGroups *[]ListVariableGroupsRespPipelineVariableGroups `json:"pipeline_variable_groups,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单页条数·
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 总条目数量。 **取值范围**： 大于等于0。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVariableGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariableGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListVariableGroupsResponse", string(data)}, " ")
}
