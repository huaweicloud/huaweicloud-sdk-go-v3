package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareCountByTagsRequest Request Object
type SearchResourceShareCountByTagsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *ResourceSharesByTagsReqBody `json:"body,omitempty"`
}

func (o SearchResourceShareCountByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareCountByTagsRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceShareCountByTagsRequest", string(data)}, " ")
}
