package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceRequest Request Object
type TagResourceRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 根、组织单元、账号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o TagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceRequest struct{}"
	}

	return strings.Join([]string{"TagResourceRequest", string(data)}, " ")
}
