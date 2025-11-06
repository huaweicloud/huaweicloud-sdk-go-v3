package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtFineGrainedSnapshotDetail **参数解释**： 细粒度备份详情。 **取值范围**： 不涉及。
type ExtFineGrainedSnapshotDetail struct {

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**： 数据库模式列表。 **取值范围**： 不涉及。
	SchemaList *[]string `json:"schema_list,omitempty"`

	// **参数解释**： 数据库表列表。 **取值范围**： 不涉及。
	TableList *[]string `json:"table_list,omitempty"`
}

func (o ExtFineGrainedSnapshotDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtFineGrainedSnapshotDetail struct{}"
	}

	return strings.Join([]string{"ExtFineGrainedSnapshotDetail", string(data)}, " ")
}
