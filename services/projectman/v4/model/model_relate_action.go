package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelateAction struct {

	// **参数解释**： 关联行为code。 **取值范围**： 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**： 关联行为名称。 **取值范围**： 不涉及。
	ActionDisplayName *string `json:"action_display_name,omitempty"`

	// **参数解释**： 关联的对象列表。 **取值范围**： 不涉及。
	RelateObjectList *[]RelationObject `json:"relate_object_list,omitempty"`
}

func (o RelateAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelateAction struct{}"
	}

	return strings.Join([]string{"RelateAction", string(data)}, " ")
}
