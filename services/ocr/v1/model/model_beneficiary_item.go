package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BeneficiaryItem struct {
	BeneficiaryName *InsurancePolicyDetail `json:"beneficiary_name,omitempty"`

	BeneficiaryType *InsurancePolicyDetail `json:"beneficiary_type,omitempty"`

	BeneficiaryOrder *InsurancePolicyDetail `json:"beneficiary_order,omitempty"`

	BeneficiaryShare *InsurancePolicyDetail `json:"beneficiary_share,omitempty"`
}

func (o BeneficiaryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BeneficiaryItem struct{}"
	}

	return strings.Join([]string{"BeneficiaryItem", string(data)}, " ")
}
