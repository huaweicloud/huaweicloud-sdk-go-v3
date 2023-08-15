package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclRuleCountResponse Response Object
type DeleteAclRuleCountResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAclRuleCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclRuleCountResponse struct{}"
	}

	return strings.Join([]string{"DeleteAclRuleCountResponse", string(data)}, " ")
}
