package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRuleAclUsingPostResponse Response Object
type AddRuleAclUsingPostResponse struct {
	Data           *RuleIdList `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o AddRuleAclUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleAclUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddRuleAclUsingPostResponse", string(data)}, " ")
}
