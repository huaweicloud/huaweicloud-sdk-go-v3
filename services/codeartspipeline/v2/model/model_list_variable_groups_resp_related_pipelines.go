package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVariableGroupsRespRelatedPipelines struct {

	// 流水线ID
	Id *string `json:"id,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`
}

func (o ListVariableGroupsRespRelatedPipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariableGroupsRespRelatedPipelines struct{}"
	}

	return strings.Join([]string{"ListVariableGroupsRespRelatedPipelines", string(data)}, " ")
}
