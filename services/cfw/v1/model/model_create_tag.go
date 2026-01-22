package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTag struct {

	// **参数解释**： 标签键 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Value *string `json:"value,omitempty"`
}

func (o CreateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTag struct{}"
	}

	return strings.Join([]string{"CreateTag", string(data)}, " ")
}
