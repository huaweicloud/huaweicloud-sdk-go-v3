package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBudgetRecordsRequest Request Object
type ListSubCustomerBudgetRecordsRequest struct {

	// |参数名称：客户ID| |参数的约束及描述：该参数必填，范围限制：1-64|
	CustomerId string `json:"customer_id"`

	// |参数名称：云经销商ID| |参数的约束及描述：该参数非必填，范围限制：0-64，如果华为云总经销商（一级经销商）需要查询云经销商的子客户预算调整记录，必须携带该字段|
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// |参数名称：操作类别| |参数的约束及描述：该参数非必填，SETTING：设置，DELETE：解除关联关系，此参数不携带时，查询所有类型数据。此参数携带值不支持为空或者空串。|
	OperationType *string `json:"operation_type,omitempty"`

	// |参数名称：预算模式| |参数的约束及描述：该参数非必填，MONTHLY：月度预算，PACKAGE：一次性预算，此参数不携带时，查询所有类型数据。此参数携带值不支持为空或者空串。|
	BudgetType *string `json:"budget_type,omitempty"`

	// |参数名称：偏移量| |参数的约束及描述：该参数非必填，范围限制：0-2147483647，从0开始，默认值为0。|
	Offset *int32 `json:"offset,omitempty"`

	// |参数名称：每次查询的数量| |参数的约束及描述：该参数非必填，范围限制：1-100，默认值为10。|
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubCustomerBudgetRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBudgetRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBudgetRecordsRequest", string(data)}, " ")
}
