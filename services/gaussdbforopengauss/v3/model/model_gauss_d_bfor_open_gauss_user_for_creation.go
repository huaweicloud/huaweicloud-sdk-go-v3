package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GaussDBforOpenGaussUserForCreation struct {

	// 数据库用户名称。  数据库用户名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，不能和系统用户名称相同。  系统用户包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”。
	Name string `json:"name"`

	// 数据库用户密码。  取值范围：非空，密码长度在8到32个字符之间，至少包含大写字母、小写字母、数字、特殊字符~!@#%^*-_=+?,三种字符的组合，不能和数据库帐号“name”或“name”的逆序相同。  建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	Password string `json:"password"`

	// 数据库账户是否只支持登入。 取值范围：不传值或设置为false，创建数据库账号包含登入数据库、创建数据库与用户权限，设置为true，只包含登入数据库权限。
	IsLoginOnly *bool `json:"is_login_only,omitempty"`
}

func (o GaussDBforOpenGaussUserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserForCreation struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserForCreation", string(data)}, " ")
}
