package model

import (
	"encoding/json"

	"strings"
)

// 迁移用户响应体
type QueryUserResp struct {
	// 任务id

	JobId *string `json:"job_id,omitempty"`
	// 是否使用全局密码

	IsGlobalPassword *string `json:"is_global_password,omitempty"`
	// 错误码

	Message *string `json:"message,omitempty"`
	// 用户列表数据

	UserList *[]QueryUserDetailResp `json:"user_list,omitempty"`
	// 角色列表数据

	RolesList *[]QueryRoleDetailResp `json:"roles_list,omitempty"`
	// 是否成功

	IsSuccess *bool `json:"is_success,omitempty"`
}

func (o QueryUserResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryUserResp struct{}"
	}

	return strings.Join([]string{"QueryUserResp", string(data)}, " ")
}
