package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskusageEntity struct {

	// Broker名称。
	BrokerName *string `json:"broker_name,omitempty" xml:"broker_name"`

	// 磁盘容量。
	DataDiskSize *string `json:"data_disk_size,omitempty" xml:"data_disk_size"`

	// 已使用的磁盘容量。
	DataDiskUse *string `json:"data_disk_use,omitempty" xml:"data_disk_use"`

	// 剩余可用的磁盘容量。
	DataDiskFree *string `json:"data_disk_free,omitempty" xml:"data_disk_free"`

	// 消息标签。
	DataDiskUsePercentage *string `json:"data_disk_use_percentage,omitempty" xml:"data_disk_use_percentage"`

	// 消息标签。
	Status *string `json:"status,omitempty" xml:"status"`

	// topic磁盘容量使用列表。
	TopicList *[]DiskusageTopicEntity `json:"topic_list,omitempty" xml:"topic_list"`
}

func (o DiskusageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskusageEntity struct{}"
	}

	return strings.Join([]string{"DiskusageEntity", string(data)}, " ")
}
