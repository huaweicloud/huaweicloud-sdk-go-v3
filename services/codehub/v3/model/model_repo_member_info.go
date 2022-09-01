package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoMemberInfo struct {

	// 用户的租户ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 用户的租户名称
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 添加的用户ID
	Id string `json:"id" xml:"id"`

	// 添加的用户名
	Name string `json:"name" xml:"name"`

	// 添加的用户权限，取值范围：30->普通成员，40->管理员
	Role int32 `json:"role" xml:"role"`
}

func (o RepoMemberInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoMemberInfo struct{}"
	}

	return strings.Join([]string{"RepoMemberInfo", string(data)}, " ")
}
