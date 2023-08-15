package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSqlTemplatesRequestBody struct {

	// 更新后SQL模板文本。
	Sql *string `json:"sql,omitempty"`

	// 更新后SQL模板名称，该名称在当前工程下必须唯一。
	SqlName *string `json:"sql_name,omitempty"`

	// SQL模板的描述信息，可以为空。
	Description *string `json:"description,omitempty"`

	// 分组名称。
	Group *string `json:"group,omitempty"`
}

func (o UpdateSqlTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSqlTemplatesRequestBody", string(data)}, " ")
}
