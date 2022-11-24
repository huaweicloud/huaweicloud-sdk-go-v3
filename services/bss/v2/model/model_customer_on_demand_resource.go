package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerOnDemandResource struct {

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode *string `json:"region_code,omitempty"`

	// 所属可用区的编码。
	AvailabilityZoneCode *string `json:"availability_zone_code,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源实例名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 资源状态： 1：正常（已开通）2：宽限期3：冻结中4：变更中5：正在关闭6：已关闭
	Status *int32 `json:"status,omitempty"`

	// 云服务产品的资源规格。如果是VM的资源规格，则需要在规格后面添加“.win”或“.linux”，例如“s2.small.1.linux”。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 按需资源的容量大小。 格式如：\"resourceInfo\": \"{\\\"specSize\\\":40.0}\"
	ResourceInfo *string `json:"resource_info,omitempty"`

	// 产品规格描述。例如： 虚拟机：“通用计算增强型|c6.2xlarge.4|8vCPUs|32GB|linux”硬盘：“云硬盘_SATA_LXH01|40.0GB”
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
}

func (o CustomerOnDemandResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerOnDemandResource struct{}"
	}

	return strings.Join([]string{"CustomerOnDemandResource", string(data)}, " ")
}
