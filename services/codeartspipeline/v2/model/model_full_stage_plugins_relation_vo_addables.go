package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoAddables struct {

	// **参数解释**： 额外属性1。 **取值范围**： 不涉及。
	AdditionalProp1 *bool `json:"additionalProp1,omitempty"`

	// **参数解释**： 额外属性2。 **取值范围**： 不涉及。
	AdditionalProp2 *bool `json:"additionalProp2,omitempty"`

	// **参数解释**： 额外属性3。 **取值范围**： 不涉及。
	AdditionalProp3 *bool `json:"additionalProp3,omitempty"`
}

func (o FullStagePluginsRelationVoAddables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoAddables struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoAddables", string(data)}, " ")
}
