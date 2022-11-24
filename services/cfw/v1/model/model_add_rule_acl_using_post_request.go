package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddRuleAclUsingPostRequest struct {
	Body *AddRuleAclDto `json:"body,omitempty"`
}

func (o AddRuleAclUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleAclUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddRuleAclUsingPostRequest", string(data)}, " ")
}
