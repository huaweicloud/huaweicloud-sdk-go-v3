package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetPasswordRequestBody struct {
	// 数据库密码。 取值范围：长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。

	Password string `json:"password"`
}

func (o ResetPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequestBody", string(data)}, " ")
}
