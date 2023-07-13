package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteListenerTagsRequest Request Object
type BatchDeleteListenerTagsRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`

	Body *BatchDeleteListenerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsRequest", string(data)}, " ")
}
