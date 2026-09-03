package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceBackupMethodRequest Request Object
type UpdateInstanceBackupMethodRequest struct {

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceBackupMethodRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceBackupMethodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBackupMethodRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBackupMethodRequest", string(data)}, " ")
}
