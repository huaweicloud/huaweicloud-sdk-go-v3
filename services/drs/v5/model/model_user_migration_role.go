package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserMigrationRole 用户迁移信息角色列表字段。当前支持的场景： - 实时迁移场景：MongoDB迁移。
type UserMigrationRole struct {

	// 角色。
	Role string `json:"role"`
}

func (o UserMigrationRole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserMigrationRole struct{}"
	}

	return strings.Join([]string{"UserMigrationRole", string(data)}, " ")
}
