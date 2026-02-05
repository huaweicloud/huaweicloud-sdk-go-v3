package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPublishAppResponse Response Object
type CheckPublishAppResponse struct {

	// 校验结果。
	Result *bool `json:"result,omitempty"`

	ValidateRule   *ValidateRuleEnum `json:"validate_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CheckPublishAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPublishAppResponse struct{}"
	}

	return strings.Join([]string{"CheckPublishAppResponse", string(data)}, " ")
}
