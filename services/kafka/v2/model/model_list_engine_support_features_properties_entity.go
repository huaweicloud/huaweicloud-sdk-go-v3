package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineSupportFeaturesPropertiesEntity **参数解释**： 实例支持的功能属性描述。
type ListEngineSupportFeaturesPropertiesEntity struct {

	// **参数解释**： Smart Connect的最大任务数。 **取值范围**： 不涉及。
	MaxTask *string `json:"max_task,omitempty"`

	// **参数解释**： Smart Connect的最小任务数。 **取值范围**： 不涉及。
	MinTask *string `json:"min_task,omitempty"`

	// **参数解释**： Smart Connect的最大节点数。 **取值范围**： 不涉及。
	MaxNode *string `json:"max_node,omitempty"`

	// **参数解释**： Smart Connect的最小节点数。 **取值范围**： 不涉及。
	MinNode *string `json:"min_node,omitempty"`
}

func (o ListEngineSupportFeaturesPropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesPropertiesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesPropertiesEntity", string(data)}, " ")
}
