package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHandshakesRequest Request Object
type ListHandshakesRequest struct {

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListHandshakesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHandshakesRequest struct{}"
	}

	return strings.Join([]string{"ListHandshakesRequest", string(data)}, " ")
}
