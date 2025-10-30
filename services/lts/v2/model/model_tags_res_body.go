package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsResBody 标签字段信息
type TagsResBody struct {

	// **参数解释：** 标签键。 **取值范围：** 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 标签值。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TagsResBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResBody struct{}"
	}

	return strings.Join([]string{"TagsResBody", string(data)}, " ")
}
