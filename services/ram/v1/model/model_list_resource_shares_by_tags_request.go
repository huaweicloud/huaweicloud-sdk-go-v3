package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSharesByTagsRequest Request Object
type ListResourceSharesByTagsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Offset *string `json:"offset,omitempty"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *ResourceSharesByTagsReqBody `json:"body,omitempty"`
}

func (o ListResourceSharesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSharesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceSharesByTagsRequest", string(data)}, " ")
}
