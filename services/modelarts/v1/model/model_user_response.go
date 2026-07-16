package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserResponse **参数解释**：用户信息。
type UserResponse struct {

	// **参数解释**：用户ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：用户名。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	Domain *Domain `json:"domain,omitempty"`
}

func (o UserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResponse struct{}"
	}

	return strings.Join([]string{"UserResponse", string(data)}, " ")
}
