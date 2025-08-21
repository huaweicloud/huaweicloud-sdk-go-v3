package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryUserDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态。 **取值范围：** 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释：** 头像地址。 **取值范围：** 不涉及。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 所属租户名称。 **取值范围：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o RepositoryUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryUserDto struct{}"
	}

	return strings.Join([]string{"RepositoryUserDto", string(data)}, " ")
}
