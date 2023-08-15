package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSystemEventRequest Request Object
type DeleteSystemEventRequest struct {

	// 系统订阅名称
	EventId string `json:"event_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteSystemEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSystemEventRequest struct{}"
	}

	return strings.Join([]string{"DeleteSystemEventRequest", string(data)}, " ")
}
