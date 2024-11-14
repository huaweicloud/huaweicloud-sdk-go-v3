package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDbCacheRuleResponse Response Object
type UpdateDbCacheRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDbCacheRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbCacheRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbCacheRuleResponse", string(data)}, " ")
}
