package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsResponse Response Object
type ListTagsResponse struct {

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 资源总数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
