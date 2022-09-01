package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerOrderV3 struct {

	// 订单ID。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 云服务类型编码。例如OBS的云服务类型编码为“hws.service.type.obs”。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty" xml:"service_type_name"`

	// 客户订单来源类型： 1：客户2：代理3：合同4：分销商
	SourceType *int32 `json:"source_type,omitempty" xml:"source_type"`

	// 订单状态。 1：待审核2：待退款3：处理中4：已取消5：已完成6：待付款9：待确认10：待发货11：待收货12：待上门取货13：换新中
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 订单类型。 1：开通2：续订3：变更4：退订11：按需转包年/包月13：试用14：转商用15：费用调整
	OrderType *int32 `json:"order_type,omitempty" xml:"order_type"`

	// 订单优惠后金额（实付价格，不含券不含卡）。
	AmountAfterDiscount *float64 `json:"amount_after_discount,omitempty" xml:"amount_after_discount"`

	// 订单金额（官网价）。 退订订单中，该金额等于currencyAfterDiscount。
	OfficialAmount *float64 `json:"official_amount,omitempty" xml:"official_amount"`

	// 订单金额度量单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 创建时间 。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 支付时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	PaymentTime *string `json:"payment_time,omitempty" xml:"payment_time"`

	// 货币编码。
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 合同ID。
	ContractId *string `json:"contract_id,omitempty" xml:"contract_id"`

	AmountInfo *AmountInfomationV2 `json:"amount_info,omitempty" xml:"amount_info"`

	// 订单创建者名称。 如果是客户自己下单，则此处返回下单操作员的登录名称；如果是运营人员从后台下单，则此处返回“运营人员”；如果是运营系统自动触发下单，则此处返回“运营系统”。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`
}

func (o CustomerOrderV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerOrderV3 struct{}"
	}

	return strings.Join([]string{"CustomerOrderV3", string(data)}, " ")
}
