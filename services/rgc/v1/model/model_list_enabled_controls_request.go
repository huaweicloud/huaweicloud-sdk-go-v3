package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnabledControlsRequest Request Object
type ListEnabledControlsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListEnabledControlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnabledControlsRequest struct{}"
	}

	return strings.Join([]string{"ListEnabledControlsRequest", string(data)}, " ")
}
