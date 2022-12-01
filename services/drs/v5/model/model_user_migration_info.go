package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户迁移信息体。
type UserMigrationInfo struct {

	// 是否迁移用户。
	IsMigrateUser bool `json:"is_migrate_user"`

	// 是否统一重置密码。取值： - true：重置密码为统一密码。 - false：不统一重置密码。 当前支持的场景：  - 实时迁移场景：MySQL迁移。
	IsSetPassword bool `json:"is_set_password"`

	// 重置后的统一密码。统一重置密码为true时必填。 约束：密码不能为空。
	Password *string `json:"password,omitempty"`

	UserList *[]UserMigrationList `json:"user_list,omitempty"`

	RoleList *[]UserMigrationRole `json:"role_list,omitempty"`
}

func (o UserMigrationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserMigrationInfo struct{}"
	}

	return strings.Join([]string{"UserMigrationInfo", string(data)}, " ")
}
