package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoCesAlarmRequest Request Object
type ShowAutoCesAlarmRequest struct {

	// **参数解释**：  数据库引擎。  **约束限制**：  不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAutoCesAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoCesAlarmRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoCesAlarmRequest", string(data)}, " ")
}
