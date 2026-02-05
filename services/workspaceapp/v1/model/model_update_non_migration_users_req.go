package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNonMigrationUsersReq 更新不进行热点会话迁移用户请求体。
type UpdateNonMigrationUsersReq struct {
	UpdateType *UpdateUserEnum `json:"update_type"`

	// 热点时不迁移用户id列表。
	NonMigrateUsers []UserInfo `json:"non_migrate_users"`
}

func (o UpdateNonMigrationUsersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNonMigrationUsersReq struct{}"
	}

	return strings.Join([]string{"UpdateNonMigrationUsersReq", string(data)}, " ")
}
