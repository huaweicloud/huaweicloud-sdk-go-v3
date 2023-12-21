package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbsGetSpecInfo 堡垒机规格信息。
type CbsGetSpecInfo struct {

	// 云堡垒机规格名称。
	ResourceSpecCode string `json:"resource_spec_code"`

	// 云堡垒机系统盘磁盘大小，单位GB。
	EcsSystemDataSize int32 `json:"ecs_system_data_size"`

	// 云堡垒机CPU核心数。
	Cpu int32 `json:"cpu"`

	// 云堡垒机内存大小，单位GB。
	Ram int32 `json:"ram"`

	// 云堡垒机资产数量。
	Asset int32 `json:"asset"`

	// 云堡垒机最大连接数。
	Connection int32 `json:"connection"`

	// 云堡垒机规格类型。 - basic：标准版 - enhance：专业版
	Type string `json:"type"`

	// 云堡垒机数据盘大小，单位TB。
	DataDiskSize float32 `json:"data_disk_size"`
}

func (o CbsGetSpecInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbsGetSpecInfo struct{}"
	}

	return strings.Join([]string{"CbsGetSpecInfo", string(data)}, " ")
}
