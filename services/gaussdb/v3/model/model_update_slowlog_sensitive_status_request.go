package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlowlogSensitiveStatusRequest Request Object
type UpdateSlowlogSensitiveStatusRequest struct {

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateSlowlogSensitiveStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateSlowlogSensitiveStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveStatusRequest", string(data)}, " ")
}
