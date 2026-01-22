package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTag struct {

	// **参数解释**： 标签键 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Value *string `json:"value,omitempty"`

	// **参数解释**： 标签更新时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
