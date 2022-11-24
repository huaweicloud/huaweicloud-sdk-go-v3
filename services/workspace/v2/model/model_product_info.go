package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 产品id。
	ProductId *string `json:"product_id,omitempty"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 产品类型。取值为： BASE：表示产品基础套餐，套餐镜像中不包括除操作系统之外的其他商业软件，私有镜像场景只能使用此类套餐。
	Type *string `json:"type,omitempty"`

	// 产品架构，当前支持：arm、x86。
	Architecture *string `json:"architecture,omitempty"`

	// cpu。
	Cpu *string `json:"cpu,omitempty"`

	// cpu描述。
	CpuDesc *string `json:"cpu_desc,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`

	// 是否是GPU类型的规格。
	IsGpu *bool `json:"is_gpu,omitempty"`

	// 系统盘类型。
	SystemDiskType *string `json:"system_disk_type,omitempty"`

	// 系统盘大小，单位GB。
	SystemDiskSize *string `json:"system_disk_size,omitempty"`

	// GPU描述。
	GpuDesc *string `json:"gpu_desc,omitempty"`

	// 话单开关，默认on,on-出话单模式,off-关话单模式,只支持反序列化，不支持序列化，不在接口中展示。
	BillSwitch *string `json:"bill_switch,omitempty"`

	// 产品描述。
	Descriptions *string `json:"descriptions,omitempty"`

	// 周期套餐标识。0表示包周期，1表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 套餐计费是否包含了数据盘,off-不包含。
	ContainDataDisk *bool `json:"contain_data_disk,omitempty"`

	// 资源规格。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务编码。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 磁盘产品类型。
	VolumeProductType *string `json:"volume_product_type,omitempty"`

	// 该产品套餐支持的专有域id（domainId）。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 产品状态，normal：正常、sellout：售空、abandon：下线。
	Status *string `json:"status,omitempty"`

	// 专属主机的子产品。
	SubProductList *[]string `json:"sub_product_list,omitempty"`

	// 套餐系列。
	PackageType *string `json:"package_type,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
