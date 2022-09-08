package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModSubCustomerBudgetReq struct {

	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 调整的目标金额。 单位：元。精确至小数点后2位。
	BudgetAmount float64 `json:"budget_amount"`

	// 是否在设置客户预算的同时解除账号冻结： 0：否1：是 默认值为0。
	CancelPartnerFrozen *string `json:"cancel_partner_frozen,omitempty"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ModSubCustomerBudgetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModSubCustomerBudgetReq struct{}"
	}

	return strings.Join([]string{"ModSubCustomerBudgetReq", string(data)}, " ")
}
