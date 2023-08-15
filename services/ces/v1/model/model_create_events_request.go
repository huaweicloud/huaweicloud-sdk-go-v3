package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventsRequest Request Object
type CreateEventsRequest struct {

	// 上报自定义事件。请求参数。
	Body *[]EventItem `json:"body,omitempty"`
}

func (o CreateEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateEventsRequest", string(data)}, " ")
}
