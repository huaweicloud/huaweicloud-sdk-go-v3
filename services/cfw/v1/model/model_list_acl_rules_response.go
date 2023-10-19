package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclRulesResponse Response Object
type ListAclRulesResponse struct {
	Data           *RuleAclListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAclRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAclRulesResponse", string(data)}, " ")
}
