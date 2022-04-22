package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUsersResponse struct {

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
	IsSuccess      *bool `json:"is_success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
