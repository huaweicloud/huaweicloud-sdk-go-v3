package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventsResponse Response Object
type CreateEventsResponse struct {

	// 事件响应体
	Body           *[]Event `json:"body,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsResponse struct{}"
	}

	return strings.Join([]string{"CreateEventsResponse", string(data)}, " ")
}
