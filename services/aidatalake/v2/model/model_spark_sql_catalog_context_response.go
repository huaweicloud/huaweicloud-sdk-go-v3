package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlCatalogContextResponse **参数解释**：Catalog上下文信息，用于指定作业使用的数据目录和数据库。
type SparkSqlCatalogContextResponse struct {

	// **参数解释**：Catalog名称，用于指定作业使用的数据目录。 **取值范围**：长度不超过128个字符。
	CatalogName *string `json:"catalog_name,omitempty"`

	// **参数解释**：默认数据库名称，用于指定作业使用的数据库。 **取值范围**：长度不超过128个字符。
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o SparkSqlCatalogContextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlCatalogContextResponse struct{}"
	}

	return strings.Join([]string{"SparkSqlCatalogContextResponse", string(data)}, " ")
}
