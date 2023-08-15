package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskEventsResponse Response Object
type ShowTaskEventsResponse struct {

	// 事件条数
	Count *int32 `json:"count,omitempty"`

	// 事件列表
	Events         *[]EventRsp `json:"events,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTaskEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskEventsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskEventsResponse", string(data)}, " ")
}
