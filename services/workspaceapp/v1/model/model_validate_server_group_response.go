package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateServerGroupResponse Response Object
type ValidateServerGroupResponse struct {

	// 校验结果。
	Result *bool `json:"result,omitempty"`

	ValidateRule   *ValidateRuleEnum `json:"validate_rule,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ValidateServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateServerGroupResponse struct{}"
	}

	return strings.Join([]string{"ValidateServerGroupResponse", string(data)}, " ")
}
