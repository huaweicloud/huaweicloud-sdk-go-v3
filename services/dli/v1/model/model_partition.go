package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Partition struct {

	// 分区名
	PartitionName string `json:"partition_name"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 最后改动时间
	LastAccessTime int64 `json:"last_access_time"`

	// 路径，外表显示，内表不显示
	Locations *[]string `json:"locations,omitempty"`

	// 最后一个ddl语句执行时间，时间戳单位：秒
	LastDdlTime *int64 `json:"last_ddl_time,omitempty"`

	// 该分区数据总行数
	NumRows *int64 `json:"num_rows,omitempty"`

	// 分区文件数
	NumFiles *int64 `json:"num_files,omitempty"`

	// 该分区总的数据大小（单位：字节）
	TotalSize *int64 `json:"total_size,omitempty"`
}

func (o Partition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Partition struct{}"
	}

	return strings.Join([]string{"Partition", string(data)}, " ")
}
