package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelServiceTasksResponse Response Object
type ListModelServiceTasksResponse struct {

	// **参数解释**： 供应商模型列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Tasks *[]ModelServiceTaskRsp `json:"tasks,omitempty"`

	// **参数解释**： 供应商模型个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListModelServiceTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelServiceTasksResponse struct{}"
	}

	return strings.Join([]string{"ListModelServiceTasksResponse", string(data)}, " ")
}
