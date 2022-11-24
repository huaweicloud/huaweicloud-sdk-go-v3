package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRuleAclUsingPutResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRuleAclUsingPutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclUsingPutResponse struct{}"
	}

	return strings.Join([]string{"ListRuleAclUsingPutResponse", string(data)}, " ")
}
