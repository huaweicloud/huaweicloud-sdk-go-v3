package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorsItems struct {

	// 产品规格编码。
	SpecCode *string `json:"spec_code,omitempty" xml:"spec_code"`

	// 云服务类型编码。
	CloudServiceTypeCode *string `json:"cloud_service_type_code,omitempty" xml:"cloud_service_type_code"`

	// 云资源类型编码。
	CloudResourceTypeCode *string `json:"cloud_resource_type_code,omitempty" xml:"cloud_resource_type_code"`

	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 - ha_rw_split： 表示读写分离实例
	CacheMode *string `json:"cache_mode,omitempty" xml:"cache_mode"`

	// 缓存引擎类型。
	Engine *string `json:"engine,omitempty" xml:"engine"`

	// 缓存版本，当缓存引擎为Redis时，取值为3.0、4.0或5.0。
	EngineVersion *string `json:"engine_version,omitempty" xml:"engine_version"`

	// Redis缓存实例的产品类型。取值当前仅支持： generic：标准类型
	ProductType *string `json:"product_type,omitempty" xml:"product_type"`

	// CPU架构类型。取值范围如下： - x86_64：X86架构 - aarch64: ARM架构
	CpuType *string `json:"cpu_type,omitempty" xml:"cpu_type"`

	// 存储类型，取值当前仅支持： DRAM:内存存储
	StorageType *string `json:"storage_type,omitempty" xml:"storage_type"`

	// 缓存容量（G Byte）。
	Capacity *[]string `json:"capacity,omitempty" xml:"capacity"`

	// 计费模式，取值范围如下： - Hourly：按需计费 - Monthly: 包月计费 - Yearly: 包周期计费
	BillingMode *[]string `json:"billing_mode,omitempty" xml:"billing_mode"`

	// 租户侧IP数量。
	TenantIpCount *int32 `json:"tenant_ip_count,omitempty" xml:"tenant_ip_count"`

	// 定价类型，取值如下： - tier: 阶梯定价，一个规格对应多个容量 - normal: 规格和容量一一对应
	PricingType *string `json:"pricing_type,omitempty" xml:"pricing_type"`

	// 是否支持专属云。
	IsDec *bool `json:"is_dec,omitempty" xml:"is_dec"`

	// 规格的其他信息。
	Attrs *[]AttrsObject `json:"attrs,omitempty" xml:"attrs"`

	// 有资源的可用区。
	FlavorsAvailableZones *[]FlavorAzObject `json:"flavors_available_zones,omitempty" xml:"flavors_available_zones"`
}

func (o FlavorsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorsItems struct{}"
	}

	return strings.Join([]string{"FlavorsItems", string(data)}, " ")
}
