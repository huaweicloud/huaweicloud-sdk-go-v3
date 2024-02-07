package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthFilterTagsRequest Request Object
type ListInternetBandwidthFilterTagsRequest struct {

	// 每页条数
	Limit *[]int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *[]int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInternetBandwidthFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthFilterTagsRequest", string(data)}, " ")
}
