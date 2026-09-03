package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupUsageAlarmConfigRequest Request Object
type UpdateBackupUsageAlarmConfigRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdateBackupUsageAlarmConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateBackupUsageAlarmConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupUsageAlarmConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateBackupUsageAlarmConfigRequest", string(data)}, " ")
}
