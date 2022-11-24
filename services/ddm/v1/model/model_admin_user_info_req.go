package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 管理员账号信息请求参数。
type AdminUserInfoReq struct {

	// 管理员账号用户名。 - 长度为1-32个字符。 - 必须以字母开头。 - 可以包含字母，数字、下划线，不能包含其它特殊字符。
	Name string `json:"name"`

	// 管理员账号密码。 - 长度为8~32位。 - 必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	Password string `json:"password"`
}

func (o AdminUserInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdminUserInfoReq struct{}"
	}

	return strings.Join([]string{"AdminUserInfoReq", string(data)}, " ")
}
