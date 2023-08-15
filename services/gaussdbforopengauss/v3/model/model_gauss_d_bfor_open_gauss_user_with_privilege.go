package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussUserWithPrivilege 用户及其权限。
type GaussDBforOpenGaussUserWithPrivilege struct {

	// 数据库帐号名称。  数据库帐号名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，不能和系统用户名称相同且帐号名称必须存在。  系统用户包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”。
	Name string `json:"name"`

	// 数据库帐号权限。 - true：只读。 - false：可读可写。
	Readonly bool `json:"readonly"`

	// schema名称。  schema名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，不能和GaussDB 模板库重名，且schema名称必须存在。  GaussDB 模板库包括postgres， template0 ，template1, public，information_schema。
	SchemaName string `json:"schema_name"`
}

func (o GaussDBforOpenGaussUserWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserWithPrivilege struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserWithPrivilege", string(data)}, " ")
}
