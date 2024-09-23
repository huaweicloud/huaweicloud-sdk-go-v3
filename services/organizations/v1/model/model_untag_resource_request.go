package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagResourceRequest Request Object
type UntagResourceRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 根、组织单元、账号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *UntagResourceReqBody `json:"body,omitempty"`
}

func (o UntagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagResourceRequest struct{}"
	}

	return strings.Join([]string{"UntagResourceRequest", string(data)}, " ")
}
