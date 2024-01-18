package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagsResponse Response Object
type ListResourcesByTagsResponse struct {
	Resources *[]Resource `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}
