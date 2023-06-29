package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateListenerTagsRequest Request Object
type BatchCreateListenerTagsRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`

	Body *BatchCreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsRequest", string(data)}, " ")
}
