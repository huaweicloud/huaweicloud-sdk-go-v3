package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Creator **参数解释**：资源所有者信息。
type Creator struct {

	// **参数解释**：租户名称。 **取值范围**：支持中文字符、英文字母、数字及特殊字符，长度为1~64个字符。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：租户ID。 **取值范围**：支持英文字母、数字及中划线，长度为1~64个字符。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：用户名称。 **取值范围**：支持英文字母、数字、下划线的组合，长度为1~64个字符。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：用户ID。 **取值范围**：支持英文字母、数字及中划线，长度为1~64个字符。
	UserId *string `json:"user_id,omitempty"`
}

func (o Creator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Creator struct{}"
	}

	return strings.Join([]string{"Creator", string(data)}, " ")
}
