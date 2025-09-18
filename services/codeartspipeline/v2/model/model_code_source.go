package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeSource struct {

	// **参数解释**： 流水线源类型，目前支持“code”、“artifact”等代码源类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	Params *CodeSourceParams `json:"params,omitempty"`
}

func (o CodeSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSource struct{}"
	}

	return strings.Join([]string{"CodeSource", string(data)}, " ")
}
