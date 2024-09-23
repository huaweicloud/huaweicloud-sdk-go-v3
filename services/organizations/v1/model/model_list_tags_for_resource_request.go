package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsForResourceRequest Request Object
type ListTagsForResourceRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 根、组织单元、账号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListTagsForResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceRequest struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceRequest", string(data)}, " ")
}
