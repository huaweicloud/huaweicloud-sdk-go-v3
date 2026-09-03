package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyNewBackupEncryptRequestBody 设置备份加密开关请求体
type ModifyNewBackupEncryptRequestBody struct {

	// **参数解释**：  KMS密钥ID，用于备份加密。  **约束限制**：  当enabled为true时必填，当enabled为false时不需填写。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	KmsKey *string `json:"kms_key,omitempty"`

	// **参数解释**：  是否开启备份加密。  **约束限制**：  不涉及。  **取值范围**：  - true：开启备份加密 - false：关闭备份加密  **默认取值**：  不涉及。
	Enabled bool `json:"enabled"`
}

func (o ModifyNewBackupEncryptRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNewBackupEncryptRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyNewBackupEncryptRequestBody", string(data)}, " ")
}
