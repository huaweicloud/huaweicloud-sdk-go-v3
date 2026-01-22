package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportRuleAclResponse Response Object
type ExportRuleAclResponse struct {
	Data           *RuleAclResponseData `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ExportRuleAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportRuleAclResponse struct{}"
	}

	return strings.Join([]string{"ExportRuleAclResponse", string(data)}, " ")
}
