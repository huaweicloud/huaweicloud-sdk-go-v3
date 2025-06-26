package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableNameResult **参数解释**： 表名检查结果。 **取值范围**： 不涉及。
type CheckTableNameResult struct {

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**： 恢复源表信息。 **取值范围**： 不涉及。
	RestoreTableList *[]string `json:"restore_table_list,omitempty"`

	// **参数解释**： 恢复目的表信息。 **取值范围**： 不涉及。
	TargetTableList *[]string `json:"target_table_list,omitempty"`
}

func (o CheckTableNameResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableNameResult struct{}"
	}

	return strings.Join([]string{"CheckTableNameResult", string(data)}, " ")
}
