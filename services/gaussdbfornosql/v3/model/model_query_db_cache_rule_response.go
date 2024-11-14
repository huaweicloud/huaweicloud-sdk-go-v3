package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDbCacheRuleResponse struct {

	// 内存加速规则ID。
	Id string `json:"id"`

	// 内存加速规则名称。
	Name *string `json:"name,omitempty"`

	// 内存加速规则状态。 - normal,正常;  - createfail, 创建失败;
	Status *string `json:"status,omitempty"`

	// 源端数据库。
	SourceDbSchema *string `json:"source_db_schema,omitempty"`

	// 源端数据表。
	SourceDbTable *string `json:"source_db_table,omitempty"`

	// 目标数据存储类型。取值为： hash。
	StorageType *string `json:"storage_type,omitempty"`

	// 目标数据库。
	TargetDatabase *string `json:"target_database,omitempty"`

	// 映射的key使用的column列表。
	KeyColumns *[]string `json:"key_columns,omitempty"`

	// 映射的value使用的column列表。
	ValueColumns *[]string `json:"value_columns,omitempty"`

	// key的生存时间。单位:ms。不传该值，默认取-1，表示永久存储。
	Ttl *string `json:"ttl,omitempty"`

	// 映射的key分隔符。
	KeySeparator *string `json:"key_separator,omitempty"`

	// 映射的value分隔符。
	ValueSeparator *string `json:"value_separator,omitempty"`

	// 键前缀。
	KeyPrefix *string `json:"key_prefix,omitempty"`
}

func (o QueryDbCacheRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDbCacheRuleResponse struct{}"
	}

	return strings.Join([]string{"QueryDbCacheRuleResponse", string(data)}, " ")
}
