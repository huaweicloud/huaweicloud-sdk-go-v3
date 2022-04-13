package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQueueTagsResponse struct {
	// 标签列表

	Tags           *[]BatchCreateOrDeleteTagReqTags `json:"tags,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowQueueTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueTagsResponse", string(data)}, " ")
}
