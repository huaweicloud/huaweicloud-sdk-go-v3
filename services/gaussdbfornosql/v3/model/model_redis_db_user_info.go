package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户信息。
type RedisDbUserInfo struct {

	// 账号名称。  小于36个字符，以字母开头，仅包含数字、字母、中划线、下划线。
	Name string `json:"name"`

	// 账号类型。  - rwuser：管理员用户 - acluser：普通用户
	Type string `json:"type"`

	// 账号权限。  - 取值\"ReadOnly\"：账号为只读权限； - 取值\"ReadWrite\"：账号为读写权限。
	Privilege string `json:"privilege"`

	// 账号已授权的数据库名称列表。
	Databases []string `json:"databases"`
}

func (o RedisDbUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisDbUserInfo struct{}"
	}

	return strings.Join([]string{"RedisDbUserInfo", string(data)}, " ")
}
