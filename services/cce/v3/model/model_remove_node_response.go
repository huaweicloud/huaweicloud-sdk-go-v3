package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveNodeResponse Response Object
type RemoveNodeResponse struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： RemoveNodesTask
	Kind *string `json:"kind,omitempty"`

	Spec *RemoveNodesSpec `json:"spec,omitempty"`

	Status         *TaskStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o RemoveNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodeResponse struct{}"
	}

	return strings.Join([]string{"RemoveNodeResponse", string(data)}, " ")
}
