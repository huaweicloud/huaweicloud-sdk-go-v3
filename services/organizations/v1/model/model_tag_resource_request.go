package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type TagResourceRequest struct {

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
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
