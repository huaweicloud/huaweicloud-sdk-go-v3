package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DesktopPoolVolumeInfo struct {

	// 批量操作磁盘的磁盘集合id。
	Id string `json:"id"`

	// 桌面数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。 - SAS：高IO。 - SSD：超高IO。
	Type string `json:"type"`

	// iops，磁盘每秒进行读写的操作次数。
	Iops *int32 `json:"iops,omitempty"`

	// 吞吐量，磁盘每秒成功传送的数据量，即读取和写入的数据量。
	Throughput *int32 `json:"throughput,omitempty"`

	// kms密钥id。变更密钥是传入密钥id；如需删除密钥则传入空字符串；默认null，不变更密钥。
	KmsId *string `json:"kms_id,omitempty"`
}

func (o DesktopPoolVolumeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopPoolVolumeInfo struct{}"
	}

	return strings.Join([]string{"DesktopPoolVolumeInfo", string(data)}, " ")
}
