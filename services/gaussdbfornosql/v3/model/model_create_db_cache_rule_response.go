package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbCacheRuleResponse Response Object
type CreateDbCacheRuleResponse struct {

	// 内存加速规则ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbCacheRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateDbCacheRuleResponse", string(data)}, " ")
}
