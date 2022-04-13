package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePolicyRuleStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePolicyRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusResponse", string(data)}, " ")
}
