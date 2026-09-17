package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InplaceMigrateNodeResponse Response Object
type InplaceMigrateNodeResponse struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： InPlaceMigrateNodesTask
	Kind *string `json:"kind,omitempty"`

	Spec *InPlaceMigratetoNodesSpec `json:"spec,omitempty"`

	Status         *TaskStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o InplaceMigrateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InplaceMigrateNodeResponse struct{}"
	}

	return strings.Join([]string{"InplaceMigrateNodeResponse", string(data)}, " ")
}
