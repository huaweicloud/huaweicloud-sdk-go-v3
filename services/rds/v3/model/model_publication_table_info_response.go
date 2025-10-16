package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicationTableInfoResponse 发布表详情。
type PublicationTableInfoResponse struct {

	// 表名。
	TableName string `json:"table_name"`

	// 命名空间。
	Schema *string `json:"schema,omitempty"`

	// 发布的字段（未传或为空说明选择所有字段）。
	Columns *[]string `json:"columns,omitempty"`

	// 主键。
	PrimaryKey *[]string `json:"primary_key,omitempty"`

	// 筛选语句。
	FilterStatement *string `json:"filter_statement,omitempty"`

	Filter *PublicationTableFilterInfoResponse `json:"filter,omitempty"`

	ArticleProperties *ArticlePropertiesResponse `json:"article_properties,omitempty"`
}

func (o PublicationTableInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationTableInfoResponse struct{}"
	}

	return strings.Join([]string{"PublicationTableInfoResponse", string(data)}, " ")
}
