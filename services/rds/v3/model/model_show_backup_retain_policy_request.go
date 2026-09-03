package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRetainPolicyRequest Request Object
type ShowBackupRetainPolicyRequest struct {

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ShowBackupRetainPolicyRequestBody `json:"body,omitempty"`
}

func (o ShowBackupRetainPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRetainPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRetainPolicyRequest", string(data)}, " ")
}
