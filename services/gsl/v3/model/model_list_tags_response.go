package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagsResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty" xml:"limit"`

	// 页码
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 记录总数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 标签记录
	Tags           *[]CmTagVo `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int        `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
