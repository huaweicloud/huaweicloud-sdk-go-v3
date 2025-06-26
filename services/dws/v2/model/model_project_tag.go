package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectTag **参数解释**： 标签详情。 **取值范围**： 不涉及。
type ProjectTag struct {

	// **参数解释**： 键。 **取值范围**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 值。 **取值范围**： 不涉及。
	Values []string `json:"values"`
}

func (o ProjectTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTag struct{}"
	}

	return strings.Join([]string{"ProjectTag", string(data)}, " ")
}
