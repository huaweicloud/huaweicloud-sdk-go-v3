package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GaussDBforOpenGaussCreateSchemaReq struct {

	// schema名称。  schema名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和GaussDB 模板库和已存在的schema重名。 GaussDB 模板库包括postgres， template0 ，template1。  已存在的schema包括public，information_schema。
	Name string `json:"name"`

	// 数据库属主用户。  数据库属主名称在1到63个字符之间，不能以“pg”和数字开头，不能和系统用户名称相同。  系统用户包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”。
	Owner string `json:"owner"`
}

func (o GaussDBforOpenGaussCreateSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussCreateSchemaReq struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussCreateSchemaReq", string(data)}, " ")
}
