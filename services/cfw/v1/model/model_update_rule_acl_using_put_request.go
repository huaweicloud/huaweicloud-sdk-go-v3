package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRuleAclUsingPutRequest struct {

	// 规则Id
	AclRuleId string `json:"acl_rule_id"`

	Body *UpdateRuleAclDto `json:"body,omitempty"`
}

func (o UpdateRuleAclUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleAclUsingPutRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleAclUsingPutRequest", string(data)}, " ")
}
