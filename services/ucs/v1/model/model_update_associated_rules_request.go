package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAssociatedRulesRequest struct {

	// 权限策略id以及涉及到的命名空间
	RuleIDNamespaces *[]RuleIdNamespaces `json:"ruleIDNamespaces,omitempty"`
}

func (o UpdateAssociatedRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatedRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssociatedRulesRequest", string(data)}, " ")
}
