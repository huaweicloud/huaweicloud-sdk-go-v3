package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDbCacheRuleRequestBody 修改内存加速规则请求体。
type UpdateDbCacheRuleRequestBody struct {

	// 内存加速规则ID。
	DbcacheRuleId string `json:"dbcache_rule_id"`

	// 映射的value使用的column列表。
	ValueColumns []string `json:"value_columns"`

	// key的生存时间。单位:ms。不传该值，默认取2592000000，表示30天。
	Ttl *string `json:"ttl,omitempty"`

	// 映射的value分隔符。只允许一个字符。
	ValueSeparator *string `json:"value_separator,omitempty"`
}

func (o UpdateDbCacheRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbCacheRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDbCacheRuleRequestBody", string(data)}, " ")
}
