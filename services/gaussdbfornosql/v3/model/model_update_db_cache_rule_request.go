package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDbCacheRuleRequest Request Object
type UpdateDbCacheRuleRequest struct {
	Body *UpdateDbCacheRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateDbCacheRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbCacheRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbCacheRuleRequest", string(data)}, " ")
}
