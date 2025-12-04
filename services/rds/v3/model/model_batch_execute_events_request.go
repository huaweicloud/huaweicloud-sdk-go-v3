package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteEventsRequest Request Object
type BatchExecuteEventsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *OperateEventReq `json:"body,omitempty"`
}

func (o BatchExecuteEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteEventsRequest struct{}"
	}

	return strings.Join([]string{"BatchExecuteEventsRequest", string(data)}, " ")
}
