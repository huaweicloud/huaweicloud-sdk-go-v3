package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 产品类型。  - BASE：表示产品基础套餐，套餐镜像中不包括除操作系统之外的其他商业软件，私有镜像场景只能使用此类套餐。
	Type *string `json:"type,omitempty"`

	// CPU。
	Cpu *string `json:"cpu,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`

	// 产品描述。
	Descriptions *string `json:"descriptions,omitempty"`

	// 周期套餐标识，0表示包周期，1表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 产品架构。
	Architecture *string `json:"architecture,omitempty"`

	// 是否是GPU类型的规格。
	IsGpu *bool `json:"is_gpu,omitempty"`

	// 套餐类型。 - ultimate：尊享版 - enterprise：企业版 - general: 通用办公版 - workstation: 云工作站 - dedicated: 专属办公版 - solver: 解算版 - agile: 敏捷办公版
	PackageType *string `json:"package_type,omitempty"`

	// 系统盘类型。
	SystemDiskType *string `json:"system_disk_type,omitempty"`

	// 系统盘大小。
	SystemDiskSize *string `json:"system_disk_size,omitempty"`

	// 套餐计费是否包含了数据盘,off-不包含。
	ContainDataDisk *bool `json:"contain_data_disk,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务类型。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 磁盘产品类型。
	VolumeProductType *string `json:"volume_product_type,omitempty"`

	// 默认在售状态，normal代表正常，sellout代表售空，abandon代表下线。
	Status *string `json:"status,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
