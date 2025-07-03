package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowsRequest Request Object
type ListFlowsRequest struct {

	// 每页记录数  取值范围[1,100]  默认值：10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFlowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowsRequest", string(data)}, " ")
}
