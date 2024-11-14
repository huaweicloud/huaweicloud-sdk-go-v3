package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbCacheRuleRequestBody 删除内存加速规则请求体。
type DeleteDbCacheRuleRequestBody struct {

	// 内存加速规则ID。
	Id string `json:"id"`
}

func (o DeleteDbCacheRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheRuleRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheRuleRequestBody", string(data)}, " ")
}
