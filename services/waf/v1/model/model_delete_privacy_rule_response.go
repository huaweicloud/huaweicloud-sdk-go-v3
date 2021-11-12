package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePrivacyRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivacyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivacyRuleResponse", string(data)}, " ")
}
