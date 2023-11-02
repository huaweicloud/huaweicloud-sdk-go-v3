package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyEventsResponse Response Object
type ListPolicyEventsResponse struct {

	// 策略事件总数
	Count *int32 `json:"count,omitempty"`

	// 策略事件列表
	Events         *[]PolicyEventRsp `json:"events,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPolicyEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyEventsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyEventsResponse", string(data)}, " ")
}
