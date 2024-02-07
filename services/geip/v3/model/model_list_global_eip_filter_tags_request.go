package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipFilterTagsRequest Request Object
type ListGlobalEipFilterTagsRequest struct {

	// 每页条数
	Limit *[]int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *[]int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListGlobalEipFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipFilterTagsRequest", string(data)}, " ")
}
