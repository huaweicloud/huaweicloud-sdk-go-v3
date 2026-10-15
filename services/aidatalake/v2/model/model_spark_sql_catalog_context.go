package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SparkSqlCatalogContext struct {

	// **参数解释**：Catalog名称，用于指定作业使用的数据目录。可在控制台的Catalog管理页面查看，或通过查询Catalog列表接口获取。 **约束限制**：不涉及。 **取值范围**：不超过128个字符。 **默认取值**：不涉及。
	CatalogName string `json:"catalog_name"`

	// **参数解释**：默认数据库名称，用于指定作业默认操作的数据库。如果未指定，则使用Catalog的默认数据库。 **约束限制**：不涉及。 **取值范围**：不超过128个字符。 **默认取值**：不涉及。
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o SparkSqlCatalogContext) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlCatalogContext struct{}"
	}

	return strings.Join([]string{"SparkSqlCatalogContext", string(data)}, " ")
}
