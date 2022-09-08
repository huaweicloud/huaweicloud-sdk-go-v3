package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 注册数据库用户请求
type RegisterDbUserRequestBody struct {

	// 数据库用户名称
	DbUsername string `json:"db_username"`

	// 数据库用户密码
	DbUserPassword string `json:"db_user_password"`

	// 数据库类型，取值为MySQL
	DatastoreType string `json:"datastore_type"`
}

func (o RegisterDbUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDbUserRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterDbUserRequestBody", string(data)}, " ")
}
