package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableTarget struct {

	// **参数解释**：  数据库名。  获取方法请参见[查询冷热分离可归档表信息](https://support.huaweicloud.com/api-taurusdb/ListArchivableTables.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DbName string `json:"db_name"`

	// **参数解释**：  表名。  获取方法请参见[查询冷热分离可归档表信息](https://support.huaweicloud.com/api-taurusdb/ListArchivableTables.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  分区名。  获取方法请参见[查询冷热分离可归档表信息](https://support.huaweicloud.com/api-taurusdb/ListArchivableTables.html)。  **约束限制**：  如果指定表下存在分区，则该参数必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`
}

func (o TableTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableTarget struct{}"
	}

	return strings.Join([]string{"TableTarget", string(data)}, " ")
}
