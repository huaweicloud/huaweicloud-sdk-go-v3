package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyRoleResult struct {

	// **参数解释**: 策略名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 策略描述。 **取值范围**: 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o AgencyRoleResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyRoleResult struct{}"
	}

	return strings.Join([]string{"AgencyRoleResult", string(data)}, " ")
}
