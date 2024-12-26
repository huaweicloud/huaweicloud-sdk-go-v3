package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type OrderInstanceV2 struct {

	// 标识要开通资源的内部ID，资源开通以后生成的ID为resource_id。
	Id *string `json:"id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源实例名。
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode *string `json:"region_code,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 云服务产品的资源规格。如果是VM的资源规格，则需要在规格后面添加“.win”或“.linux”，例如“s2.small.1.linux”。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 父资源ID。
	ParentResourceId *string `json:"parent_resource_id,omitempty"`

	// 是否是主资源。 0：非主资源1：主资源
	IsMainResource *int32 `json:"is_main_resource,omitempty"`

	// 资源状态。状态说明请参见资源状态说明。 2：使用中3：已关闭（页面不展示这个状态）4：已冻结5：已过期
	Status *int32 `json:"status,omitempty"`

	// 资源生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 资源过期时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 资源到期后的扣费策略： 0：到期进入宽限期1：到期转按需2：到期后自动删除（从生效中直接删除）3：到期后自动续费4：到期后冻结5：到期后删除（从保留期删除）  说明： 只有“3”表示该资源是自动续订，其他情况下，都是非自动续订下的到期策略。
	ExpirePolicy *int32 `json:"expire_policy,omitempty"`

	// 产品规格描述
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// 线性大小
	SpecSize *decimal.Decimal `json:"spec_size,omitempty"`

	// 线性大小单位
	SpecSizeMeasureId *int32 `json:"spec_size_measure_id,omitempty"`

	// |参数名称：资源更新时间。| |参数约束及描述：资源更新时间。UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ，如“2019-12-25T07:32:04Z”。|
	UpdateTime *string `json:"update_time,omitempty"`

	EnterpriseProject *EnterpriseProject `json:"enterprise_project,omitempty"`
}

func (o OrderInstanceV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderInstanceV2 struct{}"
	}

	return strings.Join([]string{"OrderInstanceV2", string(data)}, " ")
}
