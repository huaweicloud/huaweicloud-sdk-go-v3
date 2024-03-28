package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFlinkSqlJobTemplateRequestBody struct {

	// 模板名称，长度限制：0-57个字符。
	Name *string `json:"name,omitempty"`

	// 模板描述，长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// Stream SQL语句，至少包含source，query，sink三个部分，长度限制：0-1024*1024个字符。
	SqlBody *string `json:"sql_body,omitempty"`
}

func (o UpdateFlinkSqlJobTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobTemplateRequestBody", string(data)}, " ")
}
