package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenersByTagsRequest struct {
	Body *ListListenersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListListenersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequest", string(data)}, " ")
}
