package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Partition 分区信息
type Partition struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名字
	TableName string `json:"table_name"`

	// 分区值的列表
	PartitionValues []string `json:"partition_values"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 最后访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time"`

	// 参数表
	Parameters map[string]string `json:"parameters"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor"`
}

func (o Partition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Partition struct{}"
	}

	return strings.Join([]string{"Partition", string(data)}, " ")
}
