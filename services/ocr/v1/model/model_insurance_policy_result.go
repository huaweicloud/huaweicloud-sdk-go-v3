package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type InsurancePolicyResult struct {

	// 发卡行。
	BankName *string `json:"bank_name,omitempty" xml:"bank_name"`

	BillNumber *InsurancePolicyDetail `json:"bill_number,omitempty" xml:"bill_number"`

	Company *InsurancePolicyDetail `json:"company,omitempty" xml:"company"`

	EffectiveDate *InsurancePolicyDetail `json:"effective_date,omitempty" xml:"effective_date"`

	ApplicantName *InsurancePolicyDetail `json:"applicant_name,omitempty" xml:"applicant_name"`

	ApplicantSex *InsurancePolicyDetail `json:"applicant_sex,omitempty" xml:"applicant_sex"`

	ApplicantBirthday *InsurancePolicyDetail `json:"applicant_birthday,omitempty" xml:"applicant_birthday"`

	ApplicantIdType *InsurancePolicyDetail `json:"applicant_id_type,omitempty" xml:"applicant_id_type"`

	ApplicantIdNumber *InsurancePolicyDetail `json:"applicant_id_number,omitempty" xml:"applicant_id_number"`

	// 被保人列表（第一个默认为主被保人）。
	InsurantList *[]InsurantItem `json:"insurant_list,omitempty" xml:"insurant_list"`

	// 受益人列表。
	BeneficiaryList *[]BeneficiaryItem `json:"beneficiary_list,omitempty" xml:"beneficiary_list"`

	// 保险项目信息列表。
	InsuranceList *[]InsuranceItem `json:"insurance_list,omitempty" xml:"insurance_list"`
}

func (o InsurancePolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsurancePolicyResult struct{}"
	}

	return strings.Join([]string{"InsurancePolicyResult", string(data)}, " ")
}
