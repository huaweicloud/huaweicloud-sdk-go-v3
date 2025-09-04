package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAppGroupResponse Response Object
type CheckAppGroupResponse struct {

	// 校验结果。
	Result *bool `json:"result,omitempty"`

	ValidateRule   *ValidateRuleEnum `json:"validate_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CheckAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAppGroupResponse struct{}"
	}

	return strings.Join([]string{"CheckAppGroupResponse", string(data)}, " ")
}
