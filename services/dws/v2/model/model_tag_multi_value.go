package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagMultiValue struct {

	// **参数解释**： 标签键。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Values *[]string `json:"values,omitempty"`
}

func (o TagMultiValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMultiValue struct{}"
	}

	return strings.Join([]string{"TagMultiValue", string(data)}, " ")
}
