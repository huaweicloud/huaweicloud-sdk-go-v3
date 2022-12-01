package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户迁移信息用户列表字段。当前支持的场景：  - 实时迁移场景：MySQL迁移，MongoDB迁移。
type UserMigrationList struct {

	// 用户ID。
	Id string `json:"id"`

	// 用户。
	Account string `json:"account"`

	// 是否重置该用户密码。当前支持的场景： - 实时迁移场景：MySQL迁移。
	IsSetPassword *bool `json:"is_set_password,omitempty"`

	// 重置后的密码。统一重置密码或单个用户重置密码为true时必填，约束：密码不能为空。
	Password *string `json:"password,omitempty"`
}

func (o UserMigrationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserMigrationList struct{}"
	}

	return strings.Join([]string{"UserMigrationList", string(data)}, " ")
}
