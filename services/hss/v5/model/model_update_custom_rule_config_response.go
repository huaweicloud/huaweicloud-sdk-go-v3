package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRuleConfigResponse Response Object
type UpdateCustomRuleConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCustomRuleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRuleConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateCustomRuleConfigResponse", string(data)}, " ")
}
