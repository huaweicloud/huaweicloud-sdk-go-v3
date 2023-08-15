package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPartitionInput 添加分区信息
type AddPartitionInput struct {

	// 是否跳过已存在的分区
	IfNotExist bool `json:"if_not_exist"`

	// 添加的分区信息
	Partitions []PartitionInput `json:"partitions"`
}

func (o AddPartitionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPartitionInput struct{}"
	}

	return strings.Join([]string{"AddPartitionInput", string(data)}, " ")
}
