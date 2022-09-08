package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsuranceItem struct {
	InsuranceName *InsurancePolicyDetail `json:"insurance_name,omitempty"`

	InsurancePeriod *InsurancePolicyDetail `json:"insurance_period,omitempty"`

	InsuranceAmount *InsurancePolicyDetail `json:"insurance_amount,omitempty"`

	PaymentFrequency *InsurancePolicyDetail `json:"payment_frequency,omitempty"`

	PaymentPeriod *InsurancePolicyDetail `json:"payment_period,omitempty"`

	PaymentAmount *InsurancePolicyDetail `json:"payment_amount,omitempty"`
}

func (o InsuranceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsuranceItem struct{}"
	}

	return strings.Join([]string{"InsuranceItem", string(data)}, " ")
}
