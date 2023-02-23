package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageDescriptor struct {

	// 分区列以外的所有字段
	Columns []Column `json:"columns"`

	// 存储路径
	Location *string `json:"location,omitempty"`

	// 是否启用数据压缩
	Compressed bool `json:"compressed"`

	// 输入格式
	InputFormat *string `json:"input_format,omitempty"`

	// 输出格式
	OutputFormat *string `json:"output_format,omitempty"`

	// 分桶的桶数量
	NumberOfBuckets *int32 `json:"number_of_buckets,omitempty"`

	// 分桶字段
	BucketColumns *[]string `json:"bucket_columns,omitempty"`

	// 指定表中的每个存储桶的排序顺序的列表
	SortColumns *[]Order `json:"sort_columns,omitempty"`

	SerdeInfo *SerDeInfo `json:"serde_info"`

	// 存储描述符的参数。 key最小长度为1，最大长度为255. value最大长度为4000
	Parameters map[string]string `json:"parameters"`

	SkewedInfo *SkewedInfo `json:"skewed_info,omitempty"`

	// 数据是否会存放在子目录中
	StoredAsSubDirectories *bool `json:"stored_as_sub_directories,omitempty"`
}

func (o StorageDescriptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageDescriptor struct{}"
	}

	return strings.Join([]string{"StorageDescriptor", string(data)}, " ")
}
