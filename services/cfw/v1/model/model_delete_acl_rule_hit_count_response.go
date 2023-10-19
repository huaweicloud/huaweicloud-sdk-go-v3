package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclRuleHitCountResponse Response Object
type DeleteAclRuleHitCountResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAclRuleHitCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclRuleHitCountResponse struct{}"
	}

	return strings.Join([]string{"DeleteAclRuleHitCountResponse", string(data)}, " ")
}
