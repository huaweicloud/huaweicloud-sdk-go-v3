package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlSampleTemplate 样例模板信息。
type SqlSampleTemplate struct {

	// 语言。
	Lang *string `json:"lang,omitempty"`

	// 样例模板名称。
	Name *string `json:"name,omitempty"`

	// 样例模板内容。
	Sql *string `json:"sql,omitempty"`

	// 样例模板描述。
	Description *string `json:"description,omitempty"`

	// 样例模板分组。
	Group *string `json:"group,omitempty"`
}

func (o SqlSampleTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlSampleTemplate struct{}"
	}

	return strings.Join([]string{"SqlSampleTemplate", string(data)}, " ")
}
