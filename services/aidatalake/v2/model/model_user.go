package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// User **参数解释**：租户信息内容，包含账户名称、账户ID和用户名称、用户ID。 **约束限制**：不涉及。
type User struct {

	// **参数解释**：账户名称。 **约束限制**：不涉及。 **取值范围**：长度为[0,64]个字符。 **默认取值**：不涉及。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：账号ID。 **约束限制**：不涉及。 **取值范围**：长度为[1,36]的数字、英文字符和中划线(-)的组合。 **默认取值**：不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：用户名称。 **约束限制**：不涉及。 **取值范围**：长度为[0,64]个字符。 **默认取值**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：用户ID。 **约束限制**：不涉及。 **取值范围**：长度为[1,36]的数字、英文字符和中划线(-)的组合。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
