package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryEventAuthorDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 状态。
	State *string `json:"state,omitempty"`

	// **参数解释：** 头像地址。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 中文名。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：** 昵称。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o RepositoryEventAuthorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryEventAuthorDto struct{}"
	}

	return strings.Join([]string{"RepositoryEventAuthorDto", string(data)}, " ")
}
