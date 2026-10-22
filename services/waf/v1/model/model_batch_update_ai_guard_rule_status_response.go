package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAiGuardRuleStatusResponse Response Object
type BatchUpdateAiGuardRuleStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateAiGuardRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAiGuardRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateAiGuardRuleStatusResponse", string(data)}, " ")
}
