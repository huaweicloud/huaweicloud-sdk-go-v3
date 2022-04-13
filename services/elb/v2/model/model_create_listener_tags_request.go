package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateListenerTagsRequest struct {
	// 监听器ID。

	ListenerId string `json:"listener_id"`

	Body *CreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o CreateListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequest", string(data)}, " ")
}
