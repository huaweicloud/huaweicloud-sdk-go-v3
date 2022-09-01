package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUsersResponse struct {

	// 任务id
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 是否使用全局密码
	IsGlobalPassword *string `json:"is_global_password,omitempty" xml:"is_global_password"`

	// 错误码
	Message *string `json:"message,omitempty" xml:"message"`

	// 用户列表数据
	UserList *[]QueryUserDetailResp `json:"user_list,omitempty" xml:"user_list"`

	// 角色列表数据
	RolesList *[]QueryRoleDetailResp `json:"roles_list,omitempty" xml:"roles_list"`

	// 是否成功
	IsSuccess      *bool `json:"is_success,omitempty" xml:"is_success"`
	HttpStatusCode int   `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
