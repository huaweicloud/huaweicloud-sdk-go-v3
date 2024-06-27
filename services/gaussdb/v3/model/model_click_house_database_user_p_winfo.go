package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseDatabaseUserPWinfo 修改数据库账号密码。
type ClickHouseDatabaseUserPWinfo struct {

	// 数据库账号名。
	UserName string `json:"user_name"`

	// 数据库账号密码。 - 8-32个字符 - 至少包含以下字符中的三种：大写字母、小写字母、数字和特殊字符~！@#%^*-_=+?, - 不能与用户名或倒序的用户名相同
	Password string `json:"password"`
}

func (o ClickHouseDatabaseUserPWinfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseDatabaseUserPWinfo struct{}"
	}

	return strings.Join([]string{"ClickHouseDatabaseUserPWinfo", string(data)}, " ")
}
