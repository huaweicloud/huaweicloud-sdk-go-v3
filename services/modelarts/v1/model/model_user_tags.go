package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserTags 资源标签。
type UserTags struct {

	// **参数解释**：键。不得以\"CCE-\"或\"__type_baremetal\"开头\"。 **取值范围**：不涉及。
	Key string `json:"key"`

	// **参数解释**：值。 **取值范围**：不涉及。
	Value string `json:"value"`
}

func (o UserTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTags struct{}"
	}

	return strings.Join([]string{"UserTags", string(data)}, " ")
}
