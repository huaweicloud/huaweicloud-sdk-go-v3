package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComplianceRuleResponse Response Object
type ListComplianceRuleResponse struct {
	Body           *[]BackupComplianceRule `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListComplianceRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComplianceRuleResponse struct{}"
	}

	return strings.Join([]string{"ListComplianceRuleResponse", string(data)}, " ")
}
