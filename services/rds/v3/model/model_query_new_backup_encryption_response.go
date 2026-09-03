package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryNewBackupEncryptionResponse Response Object
type QueryNewBackupEncryptionResponse struct {

	// **参数解释**：  备份加密开关状态。  **约束限制**：  不涉及。  **取值范围**：  - true：已开启备份加密 - false：未开启备份加密  **默认取值**：  不涉及。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o QueryNewBackupEncryptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNewBackupEncryptionResponse struct{}"
	}

	return strings.Join([]string{"QueryNewBackupEncryptionResponse", string(data)}, " ")
}
