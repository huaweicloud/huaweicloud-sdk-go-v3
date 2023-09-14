package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretEventRequest Request Object
type ShowSecretEventRequest struct {

	// 事件通知的名称。
	EventName string `json:"event_name"`
}

func (o ShowSecretEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretEventRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretEventRequest", string(data)}, " ")
}
