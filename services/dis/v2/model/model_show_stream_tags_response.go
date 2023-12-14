package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStreamTagsResponse Response Object
type ShowStreamTagsResponse struct {

	// 标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowStreamTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowStreamTagsResponse", string(data)}, " ")
}
