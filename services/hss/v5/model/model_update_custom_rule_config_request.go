package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRuleConfigRequest Request Object
type UpdateCustomRuleConfigRequest struct {
	Body *UpdateCustomRuleConfigRequestInfo `json:"body,omitempty"`
}

func (o UpdateCustomRuleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRuleConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomRuleConfigRequest", string(data)}, " ")
}
