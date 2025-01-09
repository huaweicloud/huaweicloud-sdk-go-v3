package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRestrictedRuleRequest Request Object
type DeleteRestrictedRuleRequest struct {
	Body *BatchOperateAppRulesReq `json:"body,omitempty"`
}

func (o DeleteRestrictedRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRestrictedRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRestrictedRuleRequest", string(data)}, " ")
}
