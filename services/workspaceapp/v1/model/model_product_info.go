package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 产品id。
	ProductId *string `json:"product_id,omitempty"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 产品类型。 - BASE：表示产品基础套餐，套餐镜像中不包括除操作系统之外的其他商业软件，私有镜像场景只能使用此类套餐。 - ADVANCED：表示产品高级套餐，套餐镜像中包括了一些商业软件。
	Type *string `json:"type,omitempty"`

	// 产品架构，当前仅支持x86。 - x86 - arm
	Architecture *string `json:"architecture,omitempty"`

	// CPU。
	Cpu *string `json:"cpu,omitempty"`

	// CPU描述。
	CpuDesc *string `json:"cpu_desc,omitempty"`

	// 内存大小，单位兆：MB。
	Memory *string `json:"memory,omitempty"`

	// 是否是GPU类型的规格。
	IsGpu *bool `json:"is_gpu,omitempty"`

	// 系统盘类型。
	SystemDiskType *string `json:"system_disk_type,omitempty"`

	// 系统盘大小。
	SystemDiskSize *string `json:"system_disk_size,omitempty"`

	// GPU描述。
	GpuDesc *string `json:"gpu_desc,omitempty"`

	// 产品描述。
	Descriptions *string `json:"descriptions,omitempty"`

	// 套餐标识。 - 1：表示包周期。 - 0：表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 套餐计费是否包含了数据盘。
	ContainDataDisk *bool `json:"contain_data_disk,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务类型。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 磁盘产品类型。
	VolumeProductType *string `json:"volume_product_type,omitempty"`

	// 套餐默认支持的最大会话数。
	Sessions *int32 `json:"sessions,omitempty"`

	// 产品套餐在销售模式下的状态，取值自ECS的cond:operation:status。 不配置时等同于normal在售状态。 * `normal` - 正常商用 * `abandon` - 下线（即不显示） * `sellout` - 售罄 * `obt` - 公测 * `obt_sellout` - 公测售罄 * `promotion` - 推荐(等同normal，也是商用)
	Status *string `json:"status,omitempty"`

	// 产品套餐在可用区的状态，配套status使用。 > - 此参数是AZ级配置，优选取此参数的值，某个AZ没有在此参数中配置时默认使用status参数的取值。 > - 配置格式“az(xx)”。()内为某个AZ的flavor状态，()内必须要填有状态，不填为无效配置。 > - 例如：套餐在某个region的az0正常商用，az1售罄，az2公测，az3正常商用，其他az显示下线，可配置为： >   - “status”设置为：“abandon” 。 >   - “cond_operation_az”设置为：“az0(normal), az1(sellout), az2(obt), az3(normal)”。  > -  说明：如果flavor在某个AZ下的状态与status配置状态不同，必须配置该参数。
	CondOperationAz *string `json:"cond_operation_az,omitempty"`

	// 专属主机的子产品。
	SubProductList *[]string `json:"sub_product_list,omitempty"`

	// 产品属于专有的domainId。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// 套餐类型： - general：表示产品通用套餐。 - dedicated：表示产品专属主机套餐。
	PackageType *string `json:"package_type,omitempty"`

	// 系列类型
	SeriesType *string `json:"series_type,omitempty"`

	// 产品套餐过期时间,产品将在改时间点后逐步下架。
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`

	// 产品套餐支持的GPU类型。
	SupportGpuType *string `json:"support_gpu_type,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
