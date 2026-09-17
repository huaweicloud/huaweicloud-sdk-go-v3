package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseUsageInfoResp struct {

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 该数据库的cpu占比
	TotalCpu *float32 `json:"total_cpu,omitempty"`

	// 该数据库的内存占比
	TotalMemory *float32 `json:"total_memory,omitempty"`
}

func (o DatabaseUsageInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseUsageInfoResp struct{}"
	}

	return strings.Join([]string{"DatabaseUsageInfoResp", string(data)}, " ")
}
