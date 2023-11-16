package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceIpRuleResponse Response Object
type UpdateInstanceIpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceIpRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceIpRuleResponse", string(data)}, " ")
}
