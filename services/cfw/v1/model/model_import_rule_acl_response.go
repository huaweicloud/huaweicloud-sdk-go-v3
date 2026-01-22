package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportRuleAclResponse Response Object
type ImportRuleAclResponse struct {
	Data           *RuleAclResponseData `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ImportRuleAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportRuleAclResponse struct{}"
	}

	return strings.Join([]string{"ImportRuleAclResponse", string(data)}, " ")
}
