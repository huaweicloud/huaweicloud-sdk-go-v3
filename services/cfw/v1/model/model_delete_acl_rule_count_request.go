package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAclRuleCountRequest struct {
	Body *ClearAccessLogRuleHitCountsDto `json:"body,omitempty"`
}

func (o DeleteAclRuleCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclRuleCountRequest struct{}"
	}

	return strings.Join([]string{"DeleteAclRuleCountRequest", string(data)}, " ")
}
