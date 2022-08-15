package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 授予数据库用户数据库权限。
type GrantDatabasePermissionRequestBody struct {

	// 数据库用户列表，列表最大长度为50。
	Users []GrantDatabasePermission `json:"users"`
}

func (o GrantDatabasePermissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantDatabasePermissionRequestBody struct{}"
	}

	return strings.Join([]string{"GrantDatabasePermissionRequestBody", string(data)}, " ")
}
