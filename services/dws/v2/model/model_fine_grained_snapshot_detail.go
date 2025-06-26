package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FineGrainedSnapshotDetail **参数解释**： 细粒度备份信息。 **取值范围**： 不涉及。
type FineGrainedSnapshotDetail struct {

	// **参数解释**： 数据库。 **取值范围**： 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**： 模式列表。 **取值范围**： 不涉及。
	SchemaList *[]string `json:"schema_list,omitempty"`

	// **参数解释**： 表集合。 **取值范围**： 不涉及。
	TableList *[]string `json:"table_list,omitempty"`
}

func (o FineGrainedSnapshotDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FineGrainedSnapshotDetail struct{}"
	}

	return strings.Join([]string{"FineGrainedSnapshotDetail", string(data)}, " ")
}
