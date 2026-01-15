package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbParam 数据库参数名、目标数据库参数值。
type DbParam struct {

	// 数据库参数名。 取值：binlog_cache_size，binlog_stmt_cache_size，bulk_insert_buffer_size，innodb_buffer_pool_size，key_buffer_size，long_query_time，query_cache_type，read_buffer_size，read_rnd_buffer_size，sort_buffer_size，sync_binlog，innodb_buffer_pool_instances
	Key string `json:"key"`

	// 目标数据库参数值。
	TargetValue string `json:"target_value"`
}

func (o DbParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbParam struct{}"
	}

	return strings.Join([]string{"DbParam", string(data)}, " ")
}
