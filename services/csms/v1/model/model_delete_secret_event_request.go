package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretEventRequest Request Object
type DeleteSecretEventRequest struct {

	// 事件通知的名称。
	EventName string `json:"event_name"`
}

func (o DeleteSecretEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretEventRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretEventRequest", string(data)}, " ")
}
