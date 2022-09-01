package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaxInfo struct {

	// 税种。 VATISSWHTGST
	TaxClass *string `json:"taxClass,omitempty" xml:"taxClass"`

	// 税率。
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate"`

	// 税种子类。 PISCOFINSCGSTSGSTIGSTISSWHTVAT
	SubTaxClass *string `json:"subTaxClass,omitempty" xml:"subTaxClass"`

	// 税金金额。 单位：美元
	TaxAmount *float64 `json:"taxAmount,omitempty" xml:"taxAmount"`
}

func (o TaxInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaxInfo struct{}"
	}

	return strings.Join([]string{"TaxInfo", string(data)}, " ")
}
