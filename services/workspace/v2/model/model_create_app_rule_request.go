package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppRuleRequest Request Object
type CreateAppRuleRequest struct {
	Body *CreateAppRuleReq `json:"body,omitempty"`
}

func (o CreateAppRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateAppRuleRequest", string(data)}, " ")
}
