package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对象权限信息
type ObjectAuthority struct {

	// 对象名称
	Name *string `json:"name,omitempty"`

	// 角色权限集合
	RoleAuthority *[]RoleAuthority `json:"role_authority,omitempty"`
}

func (o ObjectAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectAuthority struct{}"
	}

	return strings.Join([]string{"ObjectAuthority", string(data)}, " ")
}
