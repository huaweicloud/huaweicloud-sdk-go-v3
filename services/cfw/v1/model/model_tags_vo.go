package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsVo struct {

	// **参数解释**： 规则标签ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TagId *string `json:"tag_id,omitempty"`

	// **参数解释**： 规则标签键 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TagKey *string `json:"tag_key,omitempty"`

	// **参数解释**： 规则标签值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TagValue *string `json:"tag_value,omitempty"`
}

func (o TagsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsVo struct{}"
	}

	return strings.Join([]string{"TagsVo", string(data)}, " ")
}
