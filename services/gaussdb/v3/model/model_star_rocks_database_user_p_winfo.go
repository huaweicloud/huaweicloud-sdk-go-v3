package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksDatabaseUserPWinfo 修改数据库账号密码
type StarRocksDatabaseUserPWinfo struct {

	// 数据库账号名。
	UserName string `json:"user_name"`

	// 账户密码。 - 8-32个字符 - 不能与用户名或倒序的用户名相同 - 至少包含以下字符中的三种：大写字母、小写字母、数字和特殊字符~！@#%^*-_=+?,
	Password string `json:"password"`
}

func (o StarRocksDatabaseUserPWinfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksDatabaseUserPWinfo struct{}"
	}

	return strings.Join([]string{"StarRocksDatabaseUserPWinfo", string(data)}, " ")
}
