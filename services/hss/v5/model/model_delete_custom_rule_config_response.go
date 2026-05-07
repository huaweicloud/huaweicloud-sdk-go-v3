package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRuleConfigResponse Response Object
type DeleteCustomRuleConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomRuleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRuleConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomRuleConfigResponse", string(data)}, " ")
}
