package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyNewBackupEncryptResponse Response Object
type ModifyNewBackupEncryptResponse struct {

	// **参数解释**：  设置结果。  **约束限制**：  不涉及。  **取值范围**：  - SUCCESS：设置成功  **默认取值**：  不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyNewBackupEncryptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNewBackupEncryptResponse struct{}"
	}

	return strings.Join([]string{"ModifyNewBackupEncryptResponse", string(data)}, " ")
}
