package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclRuleHitStatusResponse Response Object
type ListAclRuleHitStatusResponse struct {
	Data           *HitStatusResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAclRuleHitStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclRuleHitStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAclRuleHitStatusResponse", string(data)}, " ")
}
