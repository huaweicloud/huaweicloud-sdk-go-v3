package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Instance struct {
	// 集群的可用分区

	AvailabilityZone string `json:"availability_zone"`
	// 网卡列表，最多两个网卡。请参见•nics参数说明

	Nics []Nics `json:"nics"`
	// 实例规格： - a79fd5ae-1833-448a-88e8-3ea2b913e1f6：表示cdm.small规格，2核CPU、4G内存的虚拟机。适合PoC验证和开发测试。 - fb8fe666-6734-4b11-bc6c-43d11db3c745：表示cdm.medium规格，4核CPU、8G内存的虚拟机适合单张表规模<1000万条的场景。 - 5ddb1071-c5d7-40e0-a874-8a032e81a697：表示cdm.large规格，8核CPU、16G内存的虚拟机。适合单张表规模≥1000万条的场景。 - 6ddb1072-c5d7-40e0-a874-8a032e81a698：表示cdm.xlarge规格，16核CPU、32G内存的虚拟机。需要10GE高速带宽进行TB以上的数据量迁移时使用

	FlavorRef string `json:"flavorRef"`
	// 节点类型，当前只有“cdm”一种类型

	Type string `json:"type"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}
