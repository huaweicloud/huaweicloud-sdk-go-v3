package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInstanceEventsResponse Response Object
type ShowTaskInstanceEventsResponse struct {

	// 事件条数
	Count *int32 `json:"count,omitempty"`

	// 事件列表
	Events         *[]EventRsp `json:"events,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTaskInstanceEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstanceEventsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInstanceEventsResponse", string(data)}, " ")
}
