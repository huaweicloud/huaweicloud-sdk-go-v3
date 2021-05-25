package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubcustomerMonthlyBillsRequest struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId *string `json:"customer_id,omitempty"`
	// 消费时间。 格式固定为YYYY-MM。 示例：2018-08

	Cycle string `json:"cycle"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 计费模式。 1：包年/包月3：按需

	ChargeMode string `json:"charge_mode"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每页个数。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 账单类型。 0：消费1：退订2：华为核销

	BillType *string `json:"bill_type,omitempty"`
	// 精英服务商ID。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubcustomerMonthlyBillsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubcustomerMonthlyBillsRequest struct{}"
	}

	return strings.Join([]string{"ListSubcustomerMonthlyBillsRequest", string(data)}, " ")
}
