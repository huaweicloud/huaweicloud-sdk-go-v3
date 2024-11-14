package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbCacheRulesResponse Response Object
type ListDbCacheRulesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 内存加速映射ID。
	DbcacheMappingId *string `json:"dbcache_mapping_id,omitempty"`

	// 内存加速规则详情。
	Rules          *[]QueryDbCacheRuleResponse `json:"rules,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListDbCacheRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"ListDbCacheRulesResponse", string(data)}, " ")
}
