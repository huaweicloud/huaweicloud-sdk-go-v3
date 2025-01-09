package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsResponse Response Object
type ListProjectTagsResponse struct {

	// 标签对象
	Tags *[]Tags `json:"tags,omitempty"`

	// 标签总数
	TotalCount     *int32 `json:"totalCount,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}
