package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveNodesTask struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： RemoveNodesTask
	Kind *string `json:"kind,omitempty"`

	Spec *RemoveNodesSpec `json:"spec"`

	Status *TaskStatus `json:"status,omitempty"`
}

func (o RemoveNodesTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodesTask struct{}"
	}

	return strings.Join([]string{"RemoveNodesTask", string(data)}, " ")
}
