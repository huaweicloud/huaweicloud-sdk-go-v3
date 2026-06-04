package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IcebergPartitionField 分区规范字段，用于定义如何从源字段生成分区字段。
type IcebergPartitionField struct {

	// 源字段的id。
	SourceId int32 `json:"source_id"`

	// 分区字段的id。
	FieldId int32 `json:"field_id"`

	// 转换函数。
	Transform string `json:"transform"`

	// 分区字段名称。
	Name string `json:"name"`
}

func (o IcebergPartitionField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergPartitionField struct{}"
	}

	return strings.Join([]string{"IcebergPartitionField", string(data)}, " ")
}
