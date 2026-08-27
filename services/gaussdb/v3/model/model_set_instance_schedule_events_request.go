package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceScheduleEventsRequest Request Object
type SetInstanceScheduleEventsRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *SetInstanceScheduleEventsRequestBody `json:"body,omitempty"`
}

func (o SetInstanceScheduleEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceScheduleEventsRequest struct{}"
	}

	return strings.Join([]string{"SetInstanceScheduleEventsRequest", string(data)}, " ")
}
