package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbCacheRuleResponse Response Object
type DeleteDbCacheRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDbCacheRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheRuleResponse", string(data)}, " ")
}
