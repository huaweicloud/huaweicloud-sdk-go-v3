package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleAclsUsingGetResponse Response Object
type ListRuleAclsUsingGetResponse struct {
	Data           *RuleAclListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListRuleAclsUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclsUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListRuleAclsUsingGetResponse", string(data)}, " ")
}
