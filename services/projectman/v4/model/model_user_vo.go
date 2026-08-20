package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserVo 用户信息详情
type UserVo struct {

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 用户短ID。
	UserNumId *int32 `json:"user_num_id,omitempty"`

	// 用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 用户所属域ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 用户昵称。
	NickName *string `json:"nick_name,omitempty"`

	// 角色ID，用户在项目中具有多个角色时用英文逗号分隔。
	RoleId *string `json:"role_id,omitempty"`

	// 用户角色名称，多个角色用英文逗号分隔。
	RoleName *string `json:"role_name,omitempty"`

	// 用户角色编码，多个角色用英文逗号分隔。
	RoleCode *string `json:"role_code,omitempty"`
}

func (o UserVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserVo struct{}"
	}

	return strings.Join([]string{"UserVo", string(data)}, " ")
}
