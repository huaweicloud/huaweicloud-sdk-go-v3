package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowArchiveRuleResponse Response Object
type ShowArchiveRuleResponse struct {
	ArchiveRule    *ArchiveRuleSummary `json:"archive_rule,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowArchiveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowArchiveRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowArchiveRuleResponse", string(data)}, " ")
}
