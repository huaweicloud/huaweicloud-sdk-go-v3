package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleAclUsingDeleteResponse Response Object
type DeleteRuleAclUsingDeleteResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleAclUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleAclUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleAclUsingDeleteResponse", string(data)}, " ")
}
