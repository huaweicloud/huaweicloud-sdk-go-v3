package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopSystemEventRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 系统订阅ID
	EventId string `json:"event_id"`
}

func (o StopSystemEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSystemEventRequest struct{}"
	}

	return strings.Join([]string{"StopSystemEventRequest", string(data)}, " ")
}
