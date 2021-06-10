package model

import (
	"encoding/json"

	"strings"
)

// 更新迁移用户请求体
type UpdateUserReq struct {
	// 任务ID

	JobId string `json:"job_id"`
	// 全局密码。

	Password *string `json:"password,omitempty"`
	// 用户迁移信息，迁移用户时必填

	List *[]UserAccountVo `json:"list,omitempty"`
	// 角色迁移信息，迁移用户时必填

	UserRoles *[]UserRoleVo `json:"user_roles,omitempty"`
	// 是否设置密码

	IsSetPassword bool `json:"is_set_password"`
	// 是否迁移用户

	IsMigrateUser bool `json:"is_migrate_user"`
}

func (o UpdateUserReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
