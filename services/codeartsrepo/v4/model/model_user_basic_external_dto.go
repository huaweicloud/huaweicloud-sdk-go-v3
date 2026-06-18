package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserBasicExternalDto struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 用户名称
	Username *string `json:"username,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 头像url
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 主页url
	WebUrl *string `json:"web_url,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 是否有相关权限。 **取值范围：** true：有权限。 false：没权限。
	HasPermission *bool `json:"has_permission,omitempty"`
}

func (o UserBasicExternalDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserBasicExternalDto struct{}"
	}

	return strings.Join([]string{"UserBasicExternalDto", string(data)}, " ")
}
