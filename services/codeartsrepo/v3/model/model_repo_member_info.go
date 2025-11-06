package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoMemberInfo struct {

	// 用户的租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 用户的租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 添加的用户ID
	Id string `json:"id"`

	// 添加的用户名
	Name string `json:"name"`

	// 添加的用户权限，取值范围：20->浏览者，30->普通成员，40->管理员
	Role int32 `json:"role"`
}

func (o RepoMemberInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoMemberInfo struct{}"
	}

	return strings.Join([]string{"RepoMemberInfo", string(data)}, " ")
}
