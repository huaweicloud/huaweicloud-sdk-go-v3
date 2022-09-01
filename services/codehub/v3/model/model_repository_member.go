package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryMember struct {

	// 仓库成员描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 邮箱地址
	Email *string `json:"email,omitempty" xml:"email"`

	// 成员是否可用
	Enabled *string `json:"enabled,omitempty" xml:"enabled"`

	// 用户名
	Name *string `json:"name,omitempty" xml:"name"`

	// 仓库用户权限，取值范围：30->普通成员，40->管理员
	Role *int32 `json:"role,omitempty" xml:"role"`

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`
}

func (o RepositoryMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryMember struct{}"
	}

	return strings.Join([]string{"RepositoryMember", string(data)}, " ")
}
