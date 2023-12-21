package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailResourceInfo 资源信息。
type InstanceDetailResourceInfo struct {

	// 云堡垒机实例规格。
	Specification string `json:"specification"`

	// 订单id。
	OrderId string `json:"order_id"`

	// 云堡垒机实例的资源id，UUID格式显示。
	ResourceId string `json:"resource_id"`

	// 云堡垒机实例数据盘大小，单位TB。
	DataDiskSize float32 `json:"data_disk_size"`

	// 云堡垒机实例数据盘资源ID。
	DiskResourceId []string `json:"disk_resource_id"`
}

func (o InstanceDetailResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailResourceInfo struct{}"
	}

	return strings.Join([]string{"InstanceDetailResourceInfo", string(data)}, " ")
}
