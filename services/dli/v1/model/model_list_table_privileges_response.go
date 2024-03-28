package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablePrivilegesResponse Response Object
type ListTablePrivilegesResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 被查询的用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 权限信息。
	Privileges     *[]TablePrivilege `json:"privileges,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTablePrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablePrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ListTablePrivilegesResponse", string(data)}, " ")
}
