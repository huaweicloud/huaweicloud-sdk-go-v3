package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NewExtensionExecution **参数解释**： 插件类型。 **取值范围**： 不涉及。
type NewExtensionExecution struct {

	// **参数解释**： 插件类型。 **取值范围**： 不涉及。
	Target *string `json:"target,omitempty"`

	// **参数解释**： 插件类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： sha256。 **取值范围**： 不涉及。
	Sha256 *string `json:"sha256,omitempty"`
}

func (o NewExtensionExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionExecution struct{}"
	}

	return strings.Join([]string{"NewExtensionExecution", string(data)}, " ")
}
