package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GaussDbListDatabaseRoles struct {

	// 数据库用户/角色名。
	Name string `json:"name"`

	// 用户/角色的默认权限。
	Memberof *string `json:"memberof,omitempty"`

	// 用户/角色是否被锁。
	LockStatus *bool `json:"lock_status,omitempty"`

	Attribute *GaussDbListDatabaseRolesPriv `json:"attribute,omitempty"`
}

func (o GaussDbListDatabaseRoles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDbListDatabaseRoles struct{}"
	}

	return strings.Join([]string{"GaussDbListDatabaseRoles", string(data)}, " ")
}
