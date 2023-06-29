package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DropPartitionsInput 删除分区信息
type DropPartitionsInput struct {

	// 是否跳过不存在分区
	IfExist *bool `json:"if_exist,omitempty"`

	// 非事务表：删除分区的数据；若if_purge为真，立即释放空间。 事务表：数据保留但不可见，待数据过期统一删除。
	DeleteData *bool `json:"delete_data,omitempty"`

	// 删除分区值
	PartitionValues [][]string `json:"partition_values"`
}

func (o DropPartitionsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DropPartitionsInput struct{}"
	}

	return strings.Join([]string{"DropPartitionsInput", string(data)}, " ")
}
