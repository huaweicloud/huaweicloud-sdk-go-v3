package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BudgetRecordInfo struct {

	// |参数名称：预算金额||参数的约束及描述：单位：元，精确至小数点后2位。范围限制:0-2147483647|
	BudgetAmount *decimal.Decimal `json:"budget_amount,omitempty"`

	// |参数名称：操作类别| |参数的约束及描述：范围限制：0-10。SETTING：设置 DELETE：解除关联关系。|
	OperationType *string `json:"operation_type,omitempty"`

	// |参数名称：操作时间| |参数的约束及描述：UTC时间，格式为：yyyy-MM-ddTHH:mm:ssZ。|
	OperationTime *string `json:"operation_time,omitempty"`

	// |参数名称：操作员或系统system| |参数的约束及描述：范围限制：0-64|
	Operator *string `json:"operator,omitempty"`

	// |参数名称：预算模式| |参数的约束及描述：范围限制：0-10。MONTHLY：月度预算 PACKAGE：一次性预算|
	BudgetType *string `json:"budget_type,omitempty"`

	// |参数名称：伙伴名称| |参数的约束及描述：范围限制：0-64|
	PartnerCorpName *string `json:"partner_corp_name,omitempty"`

	// |参数名称：伙伴账号名| |参数的约束及描述：范围限制：0-256|
	PartnerAccountName *string `json:"partner_account_name,omitempty"`
}

func (o BudgetRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BudgetRecordInfo struct{}"
	}

	return strings.Join([]string{"BudgetRecordInfo", string(data)}, " ")
}
