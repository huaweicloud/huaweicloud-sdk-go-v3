package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipSegmentFilterTagsRequest Request Object
type ListGeipSegmentFilterTagsRequest struct {

	// 每页条数
	Limit *[]int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *[]int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListGeipSegmentFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipSegmentFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGeipSegmentFilterTagsRequest", string(data)}, " ")
}
