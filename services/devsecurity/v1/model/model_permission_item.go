package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionItem struct {

	// 权限名称
	PermName *string `json:"perm_name,omitempty"`

	// 权限描述
	Desc *string `json:"desc,omitempty"`

	// 权限保护级别
	ProLevelName *string `json:"pro_level_name,omitempty"`

	// 权限类型
	PermTypeName *string `json:"perm_type_name,omitempty"`
}

func (o PermissionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionItem struct{}"
	}

	return strings.Join([]string{"PermissionItem", string(data)}, " ")
}
