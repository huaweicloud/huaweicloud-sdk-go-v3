package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupResourcePackageFlavor **参数解释**:  备份资源包规格。
type BackupResourcePackageFlavor struct {

	// **参数解释**：  备份资源包规格码。  **取值范围**：  不涉及。
	SpceCode string `json:"spce_code"`

	// **参数解释**：  备份资源包规格。  **取值范围**：  不涉及。
	Volume string `json:"volume"`
}

func (o BackupResourcePackageFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResourcePackageFlavor struct{}"
	}

	return strings.Join([]string{"BackupResourcePackageFlavor", string(data)}, " ")
}
