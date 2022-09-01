package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQueueProjectTagsResponse struct {

	// 标签列表
	Tags           *[]ShowProjectTagsRespTags `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowQueueProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueProjectTagsResponse", string(data)}, " ")
}
