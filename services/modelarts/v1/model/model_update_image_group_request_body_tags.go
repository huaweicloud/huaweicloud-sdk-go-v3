package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateImageGroupRequestBodyTags struct {

	// **参数解释**：标签键 **约束限制**：最大支持20个标签键。 **取值范围**：key值最大支持长度128 **默认取值**：null。
	Key *string `json:"key,omitempty"`

	// **参数解释**：标签值 **约束限制**：最大支持20个标签值。 **取值范围**：value值最大支持长度255 **默认取值**：null。
	Value *string `json:"value,omitempty"`
}

func (o UpdateImageGroupRequestBodyTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageGroupRequestBodyTags struct{}"
	}

	return strings.Join([]string{"UpdateImageGroupRequestBodyTags", string(data)}, " ")
}
