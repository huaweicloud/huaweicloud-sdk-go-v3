package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardRuleStatusResponse Response Object
type UpdateAiGuardRuleStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAiGuardRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardRuleStatusResponse", string(data)}, " ")
}
