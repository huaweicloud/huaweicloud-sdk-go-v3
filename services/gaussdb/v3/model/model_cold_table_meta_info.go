package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColdTableMetaInfo 冷表元信息
type ColdTableMetaInfo struct {

	// **参数解释**：  表空间ID。  **取值范围**：  不涉及。
	SpaceId *string `json:"space_id,omitempty"`

	// **参数解释**：  表ID。  **取值范围**：  不涉及。
	DdId *string `json:"dd_id,omitempty"`

	// **参数解释**：  冷表库名。  **取值范围**：  不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**：  冷表表名。  **取值范围**：  不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**：  冷表分区名。  **取值范围**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  冷表有效周期（秒）。  **取值范围**：  ≥0。
	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	// **参数解释**：  冷表已保留时间（秒）。  **取值范围**：  ≥0。
	RetainedTime *int64 `json:"retained_time,omitempty"`

	// **参数解释**：  冷表数据量大小（MB）。  **取值范围**：  ≥0。
	DataSize *float32 `json:"data_size,omitempty"`
}

func (o ColdTableMetaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColdTableMetaInfo struct{}"
	}

	return strings.Join([]string{"ColdTableMetaInfo", string(data)}, " ")
}
