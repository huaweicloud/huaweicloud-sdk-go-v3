package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSystemEventRequest Request Object
type StartSystemEventRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 系统订阅ID
	EventId string `json:"event_id"`
}

func (o StartSystemEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSystemEventRequest struct{}"
	}

	return strings.Join([]string{"StartSystemEventRequest", string(data)}, " ")
}
