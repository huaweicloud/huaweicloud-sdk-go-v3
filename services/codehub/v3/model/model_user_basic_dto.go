package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserBasicDto struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 姓名
	Name *string `json:"name,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 头像url
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 头像路径
	AvatarPath *string `json:"avatar_path,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 中文名
	NameCn *string `json:"name_cn,omitempty"`

	// 主页
	WebUrl *string `json:"web_url,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o UserBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserBasicDto struct{}"
	}

	return strings.Join([]string{"UserBasicDto", string(data)}, " ")
}
