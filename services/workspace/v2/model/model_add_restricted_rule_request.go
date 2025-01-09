package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRestrictedRuleRequest Request Object
type AddRestrictedRuleRequest struct {
	Body *BatchOperateAppRulesReq `json:"body,omitempty"`
}

func (o AddRestrictedRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRestrictedRuleRequest struct{}"
	}

	return strings.Join([]string{"AddRestrictedRuleRequest", string(data)}, " ")
}
