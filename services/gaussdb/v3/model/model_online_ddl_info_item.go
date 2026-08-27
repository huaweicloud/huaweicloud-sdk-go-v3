package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OnlineDdlInfoItem struct {

	// **参数解释**：   无锁变更的目标表。  **取值范围**：  不涉及。
	Table *string `json:"table,omitempty"`

	// **参数解释**：  无锁变更的具体执行SQL。  **取值范围**：   不涉及。
	Sql *string `json:"sql,omitempty"`
}

func (o OnlineDdlInfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnlineDdlInfoItem struct{}"
	}

	return strings.Join([]string{"OnlineDdlInfoItem", string(data)}, " ")
}
