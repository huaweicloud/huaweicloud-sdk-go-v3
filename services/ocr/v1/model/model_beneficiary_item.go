package model

import (
	"encoding/json"

	"strings"
)

type BeneficiaryItem struct {
	BeneficiaryName *InsurancePolicyDetail `json:"beneficiary_name,omitempty"`

	BeneficiaryType *InsurancePolicyDetail `json:"beneficiary_type,omitempty"`
	// 受益顺序。

	BeneficiaryOrder *interface{} `json:"beneficiary_order,omitempty"`

	BeneficiaryShare *InsurancePolicyDetail `json:"beneficiary_share,omitempty"`
}

func (o BeneficiaryItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BeneficiaryItem struct{}"
	}

	return strings.Join([]string{"BeneficiaryItem", string(data)}, " ")
}
