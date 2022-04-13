package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowQueueProjectTagsRequest struct {
}

func (o ShowQueueProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueProjectTagsRequest", string(data)}, " ")
}
