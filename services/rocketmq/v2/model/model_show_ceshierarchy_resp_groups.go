package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespGroups struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 订阅Topic信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topics *[]ShowCeshierarchyRespTopics `json:"topics,omitempty"`
}

func (o ShowCeshierarchyRespGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespGroups struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespGroups", string(data)}, " ")
}
