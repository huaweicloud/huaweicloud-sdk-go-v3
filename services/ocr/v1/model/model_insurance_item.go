package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsuranceItem struct {
	InsuranceName *InsurancePolicyDetail `json:"insurance_name,omitempty" xml:"insurance_name"`

	InsurancePeriod *InsurancePolicyDetail `json:"insurance_period,omitempty" xml:"insurance_period"`

	InsuranceAmount *InsurancePolicyDetail `json:"insurance_amount,omitempty" xml:"insurance_amount"`

	PaymentFrequency *InsurancePolicyDetail `json:"payment_frequency,omitempty" xml:"payment_frequency"`

	PaymentPeriod *InsurancePolicyDetail `json:"payment_period,omitempty" xml:"payment_period"`

	PaymentAmount *InsurancePolicyDetail `json:"payment_amount,omitempty" xml:"payment_amount"`
}

func (o InsuranceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsuranceItem struct{}"
	}

	return strings.Join([]string{"InsuranceItem", string(data)}, " ")
}
