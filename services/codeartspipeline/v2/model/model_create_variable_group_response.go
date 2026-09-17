package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableGroupResponse Response Object
type CreateVariableGroupResponse struct {

	// **参数解释**： 参数组ID。 **取值范围**： 32位字符，由数字和字母组成。
	PipelineVariableGroupId *string `json:"pipeline_variable_group_id,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o CreateVariableGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateVariableGroupResponse", string(data)}, " ")
}
