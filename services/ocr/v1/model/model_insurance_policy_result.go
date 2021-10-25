package model

import (
	"encoding/json"

	"strings"
)

//
type InsurancePolicyResult struct {
	// 发卡行。

	BankName *string `json:"bank_name,omitempty"`

	BillNumber *InsurancePolicyDetail `json:"bill_number,omitempty"`

	Company *InsurancePolicyDetail `json:"company,omitempty"`

	EffectiveDate *InsurancePolicyDetail `json:"effective_date,omitempty"`

	ApplicantName *InsurancePolicyDetail `json:"applicant_name,omitempty"`

	ApplicantSex *InsurancePolicyDetail `json:"applicant_sex,omitempty"`

	ApplicantBirthday *InsurancePolicyDetail `json:"applicant_birthday,omitempty"`

	ApplicantIdType *InsurancePolicyDetail `json:"applicant_id_type,omitempty"`

	ApplicantIdNumber *InsurancePolicyDetail `json:"applicant_id_number,omitempty"`
	// 被保人列表（第一个默认为主被保人）。

	InsurantList *[]InsurantItem `json:"insurant_list,omitempty"`
	// 受益人列表。

	BeneficiaryList *[]BeneficiaryItem `json:"beneficiary_list,omitempty"`
	// 保险项目信息列表。

	InsuranceList *[]InsuranceItem `json:"insurance_list,omitempty"`
}

func (o InsurancePolicyResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InsurancePolicyResult struct{}"
	}

	return strings.Join([]string{"InsurancePolicyResult", string(data)}, " ")
}
