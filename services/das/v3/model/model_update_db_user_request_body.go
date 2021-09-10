package model

import (
	"encoding/json"

	"strings"
)

// 修改注册的数据库用户请求
type UpdateDbUserRequestBody struct {
	// 数据库用户名称

	DbUsername string `json:"db_username"`
	// 数据库用户密码

	DbUserPassword string `json:"db_user_password"`
}

func (o UpdateDbUserRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDbUserRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDbUserRequestBody", string(data)}, " ")
}
