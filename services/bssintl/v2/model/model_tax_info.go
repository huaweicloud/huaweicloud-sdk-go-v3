package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type TaxInfo struct {

	// 税种。 VATISSWHTGST
	TaxClass *string `json:"taxClass,omitempty"`

	// 税率。
	TaxRate *string `json:"taxRate,omitempty"`

	// 税种子类。 PISCOFINSCGSTSGSTIGSTISSWHTVAT
	SubTaxClass *string `json:"subTaxClass,omitempty"`

	// 税金金额。 单位：美元
	TaxAmount *decimal.Decimal `json:"taxAmount,omitempty"`

	// |参数名称：是否展示| |参数的约束及描述：是否展示。Y：展示 N：不展示|
	TaxVisibleFlag *string `json:"taxVisibleFlag,omitempty"`

	// |参数名称：是否入账| |参数的约束及描述：是否入账。Y：入账 N：不入账|
	TaxAccountingFlag *string `json:"taxAccountingFlag,omitempty"`
}

func (o TaxInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaxInfo struct{}"
	}

	return strings.Join([]string{"TaxInfo", string(data)}, " ")
}
