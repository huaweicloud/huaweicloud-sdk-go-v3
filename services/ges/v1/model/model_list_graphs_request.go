package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGraphsRequest Request Object
type ListGraphsRequest struct {

	// 本次请求的起始位置，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页资源数量的最大值，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGraphsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphsRequest struct{}"
	}

	return strings.Join([]string{"ListGraphsRequest", string(data)}, " ")
}
