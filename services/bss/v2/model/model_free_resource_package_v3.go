package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreeResourcePackageV3 struct {

	// 订购资源包产品后，系统生成的ID，是这个资源包列表的标识字段。
	OrderInstanceId *string `json:"order_instance_id,omitempty"`

	// 订单ID，如果source_type为“1：软开云赠送”，该字段为空。
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID，即资源包ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品名称，即资源包名称。
	ProductName *string `json:"product_name,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目应用范围： 0：应用所有1：应用到具体企业项目
	EnterpriseProjectScope *int32 `json:"enterprise_project_scope,omitempty"`

	// 生效时间，购买资源包的时间，格式UTC。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间，资源包到期时间，格式UTC。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 状态： 0：未生效1：生效中2：已用完3：已失效4：已退订
	Status *int32 `json:"status,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode *string `json:"region_code,omitempty"`

	// 资源包的来源类型： 0：订单1：软开云赠送
	SourceType *int32 `json:"source_type,omitempty"`

	// 套餐绑定类型： ATOMIC_PKG：原子套餐BUNDLE_PKG：组合套餐
	BundleType *string `json:"bundle_type,omitempty"`

	// 使用模式。 1：可重置表示购买的资源包能够按照一定的周期恢复使用量。例如购买一个1年的按需资源包，使用量是40G，可重置，重置周期为1个月，表示1年内每个月会给予40G的使用量。 2：不可重置表示购买的资源包的使用量不会恢复。例如购买一个1年的按需资源包，使用量是40G，不可重置，表示1年内一共给予40G的使用量。
	QuotaReuseMode *int32 `json:"quota_reuse_mode,omitempty"`

	// 资源套餐内的资源项信息（资源项ID级的详情），具体参见表3。
	FreeResources *[]FreeResourceV3 `json:"free_resources,omitempty"`
}

func (o FreeResourcePackageV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreeResourcePackageV3 struct{}"
	}

	return strings.Join([]string{"FreeResourcePackageV3", string(data)}, " ")
}
