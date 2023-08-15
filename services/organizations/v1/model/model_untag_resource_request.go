package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagResourceRequest Request Object
type UntagResourceRequest struct {

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
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
