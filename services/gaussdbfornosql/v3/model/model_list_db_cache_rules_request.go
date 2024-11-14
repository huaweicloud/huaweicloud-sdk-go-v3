package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbCacheRulesRequest Request Object
type ListDbCacheRulesRequest struct {

	// 内存加速映射ID。
	DbcacheMappingId string `json:"dbcache_mapping_id"`

	// 内存加速规则ID。
	RuleId *string `json:"rule_id,omitempty"`

	// 内存加速规则名称。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	RuleName *string `json:"rule_name,omitempty"`

	// 源端数据库名。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	SourceDbSchema *string `json:"source_db_schema,omitempty"`

	// 源端数据表名。名称以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的名称精确匹配查询。
	SourceDbTable *string `json:"source_db_table,omitempty"`

	// 索引位置，偏移量。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。 取值必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~100。不传该参数时，默认查询前100条信息。
	Limit *string `json:"limit,omitempty"`
}

func (o ListDbCacheRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbCacheRulesRequest struct{}"
	}

	return strings.Join([]string{"ListDbCacheRulesRequest", string(data)}, " ")
}
