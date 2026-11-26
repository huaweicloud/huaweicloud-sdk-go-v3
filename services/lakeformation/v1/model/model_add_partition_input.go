package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPartitionInput 添加分区信息
type AddPartitionInput struct {

	// 是否跳过已存在的分区
	IfNotExist bool `json:"if_not_exist"`

	// 添加的分区信息
	Partitions []PartitionInput `json:"partitions"`

	// 是否使用表的列作为分区中的列。默认为false
	UseTableColumn *bool `json:"use_table_column,omitempty"`
}

func (o AddPartitionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPartitionInput struct{}"
	}

	return strings.Join([]string{"AddPartitionInput", string(data)}, " ")
}
