package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabasePermissionRequestBody 删除数据库用户的数据库权限
type DeleteDatabasePermissionRequestBody struct {

	// 数据库用户列表，列表最大长度为50。
	Users []DeleteDatabasePermission `json:"users"`
}

func (o DeleteDatabasePermissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabasePermissionRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabasePermissionRequestBody", string(data)}, " ")
}
