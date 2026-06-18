package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagementUserDto struct {

	// **参数解释：** 成员名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。 **取值范围：** 不涉及
	Username *string `json:"username,omitempty"`

	// **参数解释：** 昵称。 **取值范围：** 不涉及
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 不涉及
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o ManagementUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagementUserDto struct{}"
	}

	return strings.Join([]string{"ManagementUserDto", string(data)}, " ")
}
