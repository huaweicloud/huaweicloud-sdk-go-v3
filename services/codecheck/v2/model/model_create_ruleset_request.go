package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRulesetRequest struct {
	// 设置媒体类型和编码格式

	ContentType string `json:"Content-Type"`

	Body *Ruleset `json:"body,omitempty"`
}

func (o CreateRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRulesetRequest struct{}"
	}

	return strings.Join([]string{"CreateRulesetRequest", string(data)}, " ")
}
