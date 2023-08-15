package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleAclUsingPutResponse Response Object
type UpdateRuleAclUsingPutResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRuleAclUsingPutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleAclUsingPutResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleAclUsingPutResponse", string(data)}, " ")
}
