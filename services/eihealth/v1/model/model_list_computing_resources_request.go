package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListComputingResourcesRequest struct {

	// 标签
	Label *string `json:"label,omitempty"`

	// 偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 限定查询的条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListComputingResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListComputingResourcesRequest", string(data)}, " ")
}
