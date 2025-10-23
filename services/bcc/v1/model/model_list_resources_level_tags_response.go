package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesLevelTagsResponse Response Object
type ListResourcesLevelTagsResponse struct {
	Body           *[]TagsItem `json:"body,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListResourcesLevelTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesLevelTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesLevelTagsResponse", string(data)}, " ")
}
