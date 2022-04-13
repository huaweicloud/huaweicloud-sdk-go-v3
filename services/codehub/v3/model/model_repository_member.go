package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryMember struct {
	// 仓库成员描述

	Description *string `json:"description,omitempty"`
	// 租户id

	DomainId *string `json:"domain_id,omitempty"`
	// 租户名

	DomainName *string `json:"domain_name,omitempty"`
	// 邮箱地址

	Email *string `json:"email,omitempty"`
	// 成员是否可用

	Enabled *string `json:"enabled,omitempty"`
	// 用户名

	Name *string `json:"name,omitempty"`
	// 仓库用户权限，取值范围：30->普通成员，40->管理员

	Role *int32 `json:"role,omitempty"`
	// 用户id

	UserId *string `json:"user_id,omitempty"`
}

func (o RepositoryMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryMember struct{}"
	}

	return strings.Join([]string{"RepositoryMember", string(data)}, " ")
}
