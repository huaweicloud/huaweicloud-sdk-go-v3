package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRuleAclUsingDeleteRequest struct {

	// 规则Id
	AclRuleId string `json:"acl_rule_id"`
}

func (o DeleteRuleAclUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleAclUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleAclUsingDeleteRequest", string(data)}, " ")
}
