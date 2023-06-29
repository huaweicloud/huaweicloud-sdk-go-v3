package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussDatabaseSchemaReq 创建数据库schema信息。
type GaussDBforOpenGaussDatabaseSchemaReq struct {

	// 数据库名称。  数据库名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和GaussDB 模板库重名。  GaussDB 模板库包括postgres， template0 ，template1。
	DbName string `json:"db_name"`

	// 每个元素都是与数据库相关联的schmea信息。单次请求最多支持20个元素。
	Schemas []GaussDBforOpenGaussCreateSchemaReq `json:"schemas"`
}

func (o GaussDBforOpenGaussDatabaseSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussDatabaseSchemaReq struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussDatabaseSchemaReq", string(data)}, " ")
}
