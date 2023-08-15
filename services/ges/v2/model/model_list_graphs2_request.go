package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGraphs2Request Request Object
type ListGraphs2Request struct {

	// 本次请求的起始位置，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页资源数量的最大值，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGraphs2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphs2Request struct{}"
	}

	return strings.Join([]string{"ListGraphs2Request", string(data)}, " ")
}
