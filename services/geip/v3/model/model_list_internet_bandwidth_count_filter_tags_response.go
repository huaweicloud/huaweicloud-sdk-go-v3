package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthCountFilterTagsResponse Response Object
type ListInternetBandwidthCountFilterTagsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 当前列表中资源数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInternetBandwidthCountFilterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthCountFilterTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthCountFilterTagsResponse", string(data)}, " ")
}
