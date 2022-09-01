package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BeneficiaryItem struct {
	BeneficiaryName *InsurancePolicyDetail `json:"beneficiary_name,omitempty" xml:"beneficiary_name"`

	BeneficiaryType *InsurancePolicyDetail `json:"beneficiary_type,omitempty" xml:"beneficiary_type"`

	BeneficiaryOrder *InsurancePolicyDetail `json:"beneficiary_order,omitempty" xml:"beneficiary_order"`

	BeneficiaryShare *InsurancePolicyDetail `json:"beneficiary_share,omitempty" xml:"beneficiary_share"`
}

func (o BeneficiaryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BeneficiaryItem struct{}"
	}

	return strings.Join([]string{"BeneficiaryItem", string(data)}, " ")
}
