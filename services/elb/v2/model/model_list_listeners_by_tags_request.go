package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenersByTagsRequest struct {
	Body *ListListenersByTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListListenersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequest", string(data)}, " ")
}
