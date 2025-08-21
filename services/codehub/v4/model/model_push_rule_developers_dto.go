package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PushRuleDevelopersDto struct {

	// **参数解释：** 主键ID。
	Id *interface{} `json:"id,omitempty"`

	// **参数解释：** 用户ID。
	UserId *interface{} `json:"user_id,omitempty"`

	// **参数解释：** 用户名。
	UserName *interface{} `json:"user_name,omitempty"`

	// **参数解释：** 昵称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o PushRuleDevelopersDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushRuleDevelopersDto struct{}"
	}

	return strings.Join([]string{"PushRuleDevelopersDto", string(data)}, " ")
}
