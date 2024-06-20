package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type OrderRefundInfoV2 struct {

	// 该记录的ID。
	Id string `json:"id"`

	// 金额。 金额为负数，表示退订金额。金额为正数，表示已消费金额或收取的退订手续费。
	Amount *decimal.Decimal `json:"amount"`

	// 金额的度量单位。 1：元
	MeasureId string `json:"measure_id"`

	// 客户账号ID。
	CustomerId string `json:"customer_id"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceTypeCode string `json:"resource_type_code"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode string `json:"service_type_code"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode string `json:"region_code"`

	// 退订金额、已消费金额或收取退订手续费对应的原订单ID。
	BaseOrderId *string `json:"base_order_id,omitempty"`
}

func (o OrderRefundInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderRefundInfoV2 struct{}"
	}

	return strings.Join([]string{"OrderRefundInfoV2", string(data)}, " ")
}
