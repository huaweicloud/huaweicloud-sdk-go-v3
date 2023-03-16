package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListReceivedHandshakesRequest struct {

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListReceivedHandshakesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReceivedHandshakesRequest struct{}"
	}

	return strings.Join([]string{"ListReceivedHandshakesRequest", string(data)}, " ")
}
