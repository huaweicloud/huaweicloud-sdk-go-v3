package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretEventRequest Request Object
type UpdateSecretEventRequest struct {

	// 事件通知名称。
	EventName string `json:"event_name"`

	Body *UpdateSecretEventRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretEventRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretEventRequest", string(data)}, " ")
}
