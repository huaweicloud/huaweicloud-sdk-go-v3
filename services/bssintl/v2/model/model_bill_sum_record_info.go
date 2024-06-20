package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BillSumRecordInfo struct {

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。  说明： 当请求消息中不传递“cloud_service_type_code”参数时，此值返回“null”。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 云服务区，该字段预留，先不使用。
	RegionCode *string `json:"region_code,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。
	CloudServiceTypeCode *string `json:"cloud_service_type_code,omitempty"`

	// 消费统计的时期。 格式为YYYY-MM。 示例：2018-05
	ConsumeTime *string `json:"consume_time,omitempty"`

	// 消费类型。 当请求消息中不传递“cloud_service_type_code”参数时，如果此值返回“0”表示此服务类型下所有的资源类型都是包年/包月计费模式，如果此值返回空字符串表示此服务类型下有资源类型为按需计费模式。当请求消息中传递“cloud_service_type_code”参数时，如果此值返回“0”表示此资源类型是包年/包月计费模式，如果此值返回“1”表示此资源类型为按需计费模式。
	PayMethod *string `json:"pay_method,omitempty"`

	// 消费的金额，即从客户账户实际扣除的金额。包含代金券支付的金额。
	ConsumeAmount *decimal.Decimal `json:"consume_amount,omitempty"`

	// 欠费金额，即从客户账户扣费的时候，客户账户金额不足，欠费的金额。
	Debt *decimal.Decimal `json:"debt,omitempty"`

	// 折扣金额。
	Discount *decimal.Decimal `json:"discount,omitempty"`

	// 金额单位。 1：元3：分 默认值为3。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 账单类型。 0：消费1：退订
	BillType *int32 `json:"bill_type,omitempty"`

	// 按不同账户消费类型和付费方式区分的支付总金额。 具体请参见表4。
	AccountDetails *[]BalanceTypePay `json:"account_details,omitempty"`

	// 折扣金额详情。 具体请参见表5。 当bill_type为1时，不返回此参数。
	DiscountDetailInfos *[]DiscountDetailInfo `json:"discount_detail_infos,omitempty"`

	// 企业项目ID。 当请求参数中传递了“enterpriseProjectId”，响应参数“bill_sums”返回以企业项目ID为维度的账单记录。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`
}

func (o BillSumRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillSumRecordInfo struct{}"
	}

	return strings.Join([]string{"BillSumRecordInfo", string(data)}, " ")
}
