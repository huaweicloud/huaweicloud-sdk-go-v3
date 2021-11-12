package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDatabaseRoleRequestBody struct {
	// 角色名称。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、中划线、下划线和点。

	RoleName string `json:"role_name"`
	// 角色所在的数据库名称。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、下划线。

	DbName string `json:"db_name"`
}

func (o DeleteDatabaseRoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleRequestBody", string(data)}, " ")
}
