package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseDatabaseUserInfo 账户信息。
type ClickHouseDatabaseUserInfo struct {

	// 数据库账户名。长度为2-32个字符，必须以小写字母开头，小写字母或数字结尾，可以包含小写字母、数字以及下划线，不能包含其它特殊字符。
	UserName string `json:"user_name"`

	// 账户密码。 - 8-32个字符 - 至少包含以下字符中的三种：大写字母、小写字母、数字和特殊字符~！@#%^*-_=+? - 不能与用户名或倒序的用户名相同
	Password string `json:"password"`

	// 数据库列表。“*”表示所有数据库。
	Databases []string `json:"databases"`

	// DML权限，默认2。 取值范围： - 1：只读权限 - 2：读取和设置权限
	Dml *int32 `json:"dml,omitempty"`
}

func (o ClickHouseDatabaseUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseDatabaseUserInfo struct{}"
	}

	return strings.Join([]string{"ClickHouseDatabaseUserInfo", string(data)}, " ")
}
