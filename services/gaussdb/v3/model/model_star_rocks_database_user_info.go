package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksDatabaseUserInfo 账户信息
type StarRocksDatabaseUserInfo struct {

	// 数据库账户名。 长度为2-32个字符，必须以小写字母开头，小写字母或数字结尾，可以包含小写字母、数字以及下划线，不能包含其它特殊字符。
	UserName string `json:"user_name"`

	// 账户密码。 - 8-32个字符 - 不能与用户名或倒序的用户名相同 - 至少包含以下字符中的三种：大写字母、小写字母、数字和特殊字符~！@#%^*-_=+?,
	Password string `json:"password"`

	// 数据库列表。
	Databases []string `json:"databases"`

	// DML权限，默认2。 取值范围： - 0：读写权限 - 1：只读权限 - 2：只读和设置权限 - 3：读写和设置权限
	Dml *int32 `json:"dml,omitempty"`

	// DDL权限，默认0。 取值范围： - 0：无DDL权限 - 1：有DDL权限
	Ddl *int32 `json:"ddl,omitempty"`
}

func (o StarRocksDatabaseUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksDatabaseUserInfo struct{}"
	}

	return strings.Join([]string{"StarRocksDatabaseUserInfo", string(data)}, " ")
}
