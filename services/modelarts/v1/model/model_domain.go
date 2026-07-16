package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Domain **参数解释**：Domain账号信息。
type Domain struct {

	// **参数解释**：账号ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：账号名。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o Domain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Domain struct{}"
	}

	return strings.Join([]string{"Domain", string(data)}, " ")
}
