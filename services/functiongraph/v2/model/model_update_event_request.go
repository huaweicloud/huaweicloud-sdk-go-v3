package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEventRequest Request Object
type UpdateEventRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 测试事件ID
	EventId string `json:"event_id"`

	Body *UpdateEventRequestBody `json:"body,omitempty"`
}

func (o UpdateEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventRequest", string(data)}, " ")
}
