package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeInfo 磁盘信息。
type VolumeInfo struct {

	// 批量操作磁盘的磁盘集合id。
	Id *string `json:"id,omitempty"`

	// 桌面数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。 - SAS：高IO。 - SSD：超高IO。
	Type string `json:"type"`

	// 磁盘容量，单位GB。
	Size int32 `json:"size"`

	// iops，云硬盘每秒进行读写的操作次数。
	Iops *int32 `json:"iops,omitempty"`

	// 吞吐量，云硬盘每秒成功传送的数据量，即读取和写入的数据量。
	Throughput *int32 `json:"throughput,omitempty"`

	// 规格。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// kms密钥Id
	KmsId *string `json:"kms_id,omitempty"`
}

func (o VolumeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeInfo struct{}"
	}

	return strings.Join([]string{"VolumeInfo", string(data)}, " ")
}
