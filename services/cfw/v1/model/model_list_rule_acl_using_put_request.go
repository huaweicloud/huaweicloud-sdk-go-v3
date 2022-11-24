package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRuleAclUsingPutRequest struct {

	// 规则id
	AclRuleId string `json:"acl_rule_id"`

	Body *OrderRuleAclDto `json:"body,omitempty"`
}

func (o ListRuleAclUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclUsingPutRequest struct{}"
	}

	return strings.Join([]string{"ListRuleAclUsingPutRequest", string(data)}, " ")
}
