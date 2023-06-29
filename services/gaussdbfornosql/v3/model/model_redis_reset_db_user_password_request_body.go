package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisResetDbUserPasswordRequestBody 重置数据库账号请求体。
type RedisResetDbUserPasswordRequestBody struct {

	// 账号名称。  小于36个字符，以字母开头，仅包含数字、字母、中划线、下划线。
	Name string `json:"name"`

	// 需重置的密码。  - 密码长度为8~32位。  - 密码需包含大写字母、小写字母、数字和特殊字符中的至少三种，支持的特殊字符为!@#$%^&*()_+-= 。
	Password string `json:"password"`
}

func (o RedisResetDbUserPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisResetDbUserPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"RedisResetDbUserPasswordRequestBody", string(data)}, " ")
}
