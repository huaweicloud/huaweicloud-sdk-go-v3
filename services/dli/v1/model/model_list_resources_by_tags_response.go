package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagsResponse Response Object
type ListResourcesByTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}
