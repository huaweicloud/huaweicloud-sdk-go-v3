package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompatibleFaultResp **参数解释**： 信息。 **取值范围**： 不涉及。
type CompatibleFaultResp struct {

	// **参数解释**： 信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**： 创建者。 **取值范围**： 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**： 详细内容。 **取值范围**： 不涉及。
	Details *string `json:"details,omitempty"`
}

func (o CompatibleFaultResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleFaultResp struct{}"
	}

	return strings.Join([]string{"CompatibleFaultResp", string(data)}, " ")
}
