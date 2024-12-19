package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseSchemaRequestBody 删除数据库schema信息。
type DeleteDatabaseSchemaRequestBody struct {

	// 数据库名称。 使用已存在的数据库名称，且不能为模板库。 模板库包括postgres， template0 ，template1，templatea，template_pdb，templatem。
	DbName string `json:"db_name"`

	// SCHEMA名称。 SCHEMA名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和模板库和已存在的SCHEMA重名。 模板库包括postgres， template0 ，template1，templatea，template_pdb，templatem。 已存在的SCHEMA包括public，information_schema。
	Schema string `json:"schema"`
}

func (o DeleteDatabaseSchemaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseSchemaRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseSchemaRequestBody", string(data)}, " ")
}
