package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlJobTemplateRequestBody struct {

	// 新增的SQL模板。
	Sql string `json:"sql"`

	// 新增SQL模板名称，该名称在当前工程下必须唯一。
	SqlName string `json:"sql_name"`

	// 新增SQL模板的描述信息，可以为空字符串。
	Description *string `json:"description,omitempty"`

	// 分组名称。
	Group *string `json:"group,omitempty"`
}

func (o CreateSqlJobTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlJobTemplateRequestBody", string(data)}, " ")
}
