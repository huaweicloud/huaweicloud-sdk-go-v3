package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomRuleConfigRequest Request Object
type CreateCustomRuleConfigRequest struct {
	Body *CreateCustomRuleConfigRequestInfo `json:"body,omitempty"`
}

func (o CreateCustomRuleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRuleConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomRuleConfigRequest", string(data)}, " ")
}
