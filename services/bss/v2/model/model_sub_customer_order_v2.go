package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubCustomerOrderV2 struct {

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 客户订单订单来源类型： 1：客户2：代理3：合同4：分销商
	SourceType *int32 `json:"source_type,omitempty"`

	// 订单状态： 1：待审核2：待退款3：处理中4：已取消5：已完成6：待支付9：待确认10：待发货11：待收货12：待上门取货13：换新中
	Status *int32 `json:"status,omitempty"`

	// 订单类型： 1：开通2：续订3：变更4：退订11：按需转包年/包月13：试用14：转商用15：费用调整
	OrderType *int32 `json:"order_type,omitempty"`

	// 订单优惠后金额（不含券不含卡的实付价格）。
	AmountAfterDiscount *float64 `json:"amount_after_discount,omitempty"`

	// 订单金额（官网价）。 退订订单中，该金额等于amount_after_discount。
	OfficialAmount *float64 `json:"official_amount,omitempty"`

	// 订单金额度量单位： 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 创建时间 。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	CreateTime *string `json:"create_time,omitempty"`

	// 支付时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	PaymentTime *string `json:"payment_time,omitempty"`

	// 货币编码。
	Currency *string `json:"currency,omitempty"`

	// 合同ID。
	ContractId *string `json:"contract_id,omitempty"`

	AmountInfo *AmountInfomationV2 `json:"amount_info,omitempty"`

	// 客户订单企业项目信息。具体请参见表 EnterpriseProject。
	EnterpriseProjects *[]EnterpriseProject `json:"enterprise_projects,omitempty"`
}

func (o SubCustomerOrderV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCustomerOrderV2 struct{}"
	}

	return strings.Join([]string{"SubCustomerOrderV2", string(data)}, " ")
}
