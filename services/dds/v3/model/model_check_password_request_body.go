package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckPasswordRequestBody struct {

	// 数据库密码。取值范围：长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。
	UserPwd string `json:"user_pwd"`

	// 数据库用户名称，默认为“rwuser”。取值范围：长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、中划线、下划线和点。
	UserName *string `json:"user_name,omitempty"`

	// 用户所在的数据库，默认为“admin”。取值范围：长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、下划线。
	DbName *string `json:"db_name,omitempty"`
}

func (o CheckPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"CheckPasswordRequestBody", string(data)}, " ")
}
