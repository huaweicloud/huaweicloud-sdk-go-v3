package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsRequest Request Object
type ListControlsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListControlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsRequest struct{}"
	}

	return strings.Join([]string{"ListControlsRequest", string(data)}, " ")
}
