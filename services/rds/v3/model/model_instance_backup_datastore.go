package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceBackupDatastore 实例引擎信息
type InstanceBackupDatastore struct {

	// **参数解释**：  数据库引擎类型  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  数据库引擎版本  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o InstanceBackupDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceBackupDatastore struct{}"
	}

	return strings.Join([]string{"InstanceBackupDatastore", string(data)}, " ")
}
