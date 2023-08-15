package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedisUserForCreation struct {

	// 账号名称。  小于36个字符，以字母开头，仅包含数字、字母、中划线、下划线。
	Name string `json:"name"`

	// - 密码长度为8~32位。  - 密码需包含大写字母、小写字母、数字和特殊字符中的至少三种，支持的特殊字符为!@#$%^&*()_+-= 。
	Password string `json:"password"`

	// 账号授权的数据库名称列表，至少指定一个数据库。
	Databases []string `json:"databases"`

	// 账号权限。  - 取值\"ReadOnly\"：账号为只读权限； - 取值\"ReadWrite\"：账号为读写权限。
	Privilege string `json:"privilege"`
}

func (o RedisUserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisUserForCreation struct{}"
	}

	return strings.Join([]string{"RedisUserForCreation", string(data)}, " ")
}
