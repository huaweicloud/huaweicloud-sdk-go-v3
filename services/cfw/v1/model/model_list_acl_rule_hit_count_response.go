package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclRuleHitCountResponse Response Object
type ListAclRuleHitCountResponse struct {
	Data           *RuleHitCountRecords `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAclRuleHitCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclRuleHitCountResponse struct{}"
	}

	return strings.Join([]string{"ListAclRuleHitCountResponse", string(data)}, " ")
}
