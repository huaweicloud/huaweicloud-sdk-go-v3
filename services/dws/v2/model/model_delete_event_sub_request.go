package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventSubRequest Request Object
type DeleteEventSubRequest struct {

	// 事件订阅ID
	EventSubId string `json:"event_sub_id"`
}

func (o DeleteEventSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSubRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventSubRequest", string(data)}, " ")
}
