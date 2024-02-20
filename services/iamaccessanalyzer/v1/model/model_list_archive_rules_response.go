package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArchiveRulesResponse Response Object
type ListArchiveRulesResponse struct {
	ArchiveRules *[]ArchiveRuleSummary `json:"archive_rules,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListArchiveRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArchiveRulesResponse struct{}"
	}

	return strings.Join([]string{"ListArchiveRulesResponse", string(data)}, " ")
}
