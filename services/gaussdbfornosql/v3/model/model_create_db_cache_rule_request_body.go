package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDbCacheRuleRequestBody struct {

	// 内存加速映射ID。
	DbcacheMappingId string `json:"dbcache_mapping_id"`

	// 内存加速规则名称。不超过256字符，需要确保在当前映射下唯一。
	Name string `json:"name"`

	// 源端数据库。
	SourceDbSchema string `json:"source_db_schema"`

	// 源端数据表。
	SourceDbTable string `json:"source_db_table"`

	// 目标数据存储类型。 取值为： hash。
	StorageType string `json:"storage_type"`

	// 目标数据库。取值范围： [0-999]。
	TargetDatabase string `json:"target_database"`

	// 映射的key使用的column列表。
	KeyColumns []string `json:"key_columns"`

	// 映射的value使用的column列表。
	ValueColumns []string `json:"value_columns"`

	// key的生存时间。单位:ms。不传该值，默认取2592000000，表示30天。
	Ttl *string `json:"ttl,omitempty"`

	// 映射的key分隔符。只允许一个字符。
	KeySeparator string `json:"key_separator"`

	// 映射的value分隔符。只允许一个字符。
	ValueSeparator *string `json:"value_separator,omitempty"`

	// 键前缀。最长限制1024字符。需确保key_prefix和target_database在当前映射下唯一。
	KeyPrefix string `json:"key_prefix"`
}

func (o CreateDbCacheRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDbCacheRuleRequestBody", string(data)}, " ")
}
