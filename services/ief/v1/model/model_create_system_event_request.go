package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSystemEventRequest Request Object
type CreateSystemEventRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EventCreateReq `json:"body,omitempty"`
}

func (o CreateSystemEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSystemEventRequest struct{}"
	}

	return strings.Join([]string{"CreateSystemEventRequest", string(data)}, " ")
}
