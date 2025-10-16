package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicationTableInfoRequest 发布表详情。
type PublicationTableInfoRequest struct {

	// 表名。  表名长度可在1～64个字符之间，由字母、数字或下划线组成，不能包含其他特殊字符
	TableName string `json:"table_name"`

	// 命名空间。默认值dbo。
	Schema *string `json:"schema,omitempty"`

	// 发布的字段（不传或为空说明选择所有字段）。  字段名称长度可在1～64个字符之间，由字母、数字或下划线组成，不能包含其他特殊字符
	Columns *[]string `json:"columns,omitempty"`

	// 主键。
	PrimaryKey *[]string `json:"primary_key,omitempty"`

	// 筛选语句。不能包含英文单引号'。
	FilterStatement *string `json:"filter_statement,omitempty"`

	Filter *PublicationTableFilterInfoRequest `json:"filter,omitempty"`

	ArticleProperties *ArticlePropertiesRequest `json:"article_properties,omitempty"`
}

func (o PublicationTableInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationTableInfoRequest struct{}"
	}

	return strings.Join([]string{"PublicationTableInfoRequest", string(data)}, " ")
}
