package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionInput partition base object
type PartitionInput struct {

	// 分区值列表
	PartitionValues []string `json:"partition_values"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 最后访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time"`

	// 参数信息
	Parameters map[string]string `json:"parameters"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor"`
}

func (o PartitionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionInput struct{}"
	}

	return strings.Join([]string{"PartitionInput", string(data)}, " ")
}
