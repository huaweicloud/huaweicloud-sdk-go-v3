package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseInfoItem struct {

	// 数据库名称, config admin 库不做展示。
	Name *string `json:"name,omitempty"`

	// 数据库存储大小（以GB为单位）,保留两位小数。 存储大小: 实际磁盘上占用的物理空间大小，包括数据文件、日志文件、索引文件等。
	DataSize *string `json:"data_size,omitempty"`

	// 数据库逻辑大小 （以GB为单位）,保留两位小数。 逻辑大小指的是数据库中实际存储的数据大小，不包括索引大小、日志大小等。
	StorageSize *string `json:"storage_size,omitempty"`

	// 数据库中的集合数。
	CollectionNum *int32 `json:"collection_num,omitempty"`
}

func (o DatabaseInfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInfoItem struct{}"
	}

	return strings.Join([]string{"DatabaseInfoItem", string(data)}, " ")
}
