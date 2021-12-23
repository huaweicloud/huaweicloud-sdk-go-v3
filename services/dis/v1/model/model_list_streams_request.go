package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListStreamsRequest struct {
	// 单次请求返回通道列表的最大数量。  取值范围：1~100。  默认值：10。

	Limit *int32 `json:"limit,omitempty"`
	// 从该通道开始返回通道列表，返回的通道列表不包括此通道名称。  如果需要分页查询，第一页查询时不传该字段。返回结果has_more_streams为true时，进行下一页查询，start_stream_name传入第一页查询结果的最后一条通道名称。

	StartStreamName *string `json:"start_stream_name,omitempty"`
}

func (o ListStreamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListStreamsRequest", string(data)}, " ")
}
