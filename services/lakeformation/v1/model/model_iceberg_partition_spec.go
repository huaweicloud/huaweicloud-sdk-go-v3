package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IcebergPartitionSpec 定义了Iceberg表的分区规格，决定了分区表数据在查询优化时的性能。
type IcebergPartitionSpec struct {

	// 分区规范id。
	SpecId int32 `json:"spec_id"`

	// IcebergPartitionField的列表。
	Fields []IcebergPartitionField `json:"fields"`
}

func (o IcebergPartitionSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergPartitionSpec struct{}"
	}

	return strings.Join([]string{"IcebergPartitionSpec", string(data)}, " ")
}
