package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudtableSchema 转储HBase时必选，与“opentsdb_schema”二选一，表示CloudTable集群HBase数据的Schema配置。用于将通道内的JSON数据进行格式转换并导入Cloudtable的HBase表中。
type CloudtableSchema struct {

	// CloudTable集群HBase数据rowkey的Schema配置，用于将通道内的JSON数据生成HBase数据的rowkey。  取值范围：1～64。
	RowKey []RowKey `json:"row_key"`

	// CloudTable集群HBase数据列的Schema配置，用于将通道内的JSON数据生成HBase数据的列。  取值范围：1～4096。
	Columns []Column `json:"columns"`
}

func (o CloudtableSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudtableSchema struct{}"
	}

	return strings.Join([]string{"CloudtableSchema", string(data)}, " ")
}
