package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AcceptanceBillResult struct {

	// 出票日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 汇票到期日。
	DueDate *string `json:"due_date,omitempty"`

	// 票据状态。
	BillStatus *string `json:"bill_status,omitempty"`

	// 票据号码。
	BillNumber *string `json:"bill_number,omitempty"`

	// 出票人全称。
	IssuerFullName *string `json:"issuer_full_name,omitempty"`

	// 出票人账号。
	IssuerAccount *string `json:"issuer_account,omitempty"`

	// 出票人开户银行。
	IssuerBankName *string `json:"issuer_bank_name,omitempty"`

	// 出票人开户行号。
	IssuerBankNumber *string `json:"issuer_bank_number,omitempty"`

	// 收款人全称。
	PayeeFullName *string `json:"payee_full_name,omitempty"`

	// 收款人账号。
	PayeeAccount *string `json:"payee_account,omitempty"`

	// 收款人开户银行。
	PayeeBankName *string `json:"payee_bank_name,omitempty"`

	// 收款人开户行号。
	PayeeBankNumber *string `json:"payee_bank_number,omitempty"`

	// 出票保证人名称。
	IssuanceGuarantorName *string `json:"issuance_guarantor_name,omitempty"`

	// 出票保证人地址。
	IssuanceGuarantorAddress *string `json:"issuance_guarantor_address,omitempty"`

	// 出票保证人账号。
	IssuanceGuarantorAccount *string `json:"issuance_guarantor_account,omitempty"`

	// 出票保证日期。
	IssuanceGuaranteeDate *string `json:"issuance_guarantee_date,omitempty"`

	// 出票保证人开户行行号。
	IssuanceGuarantorBankNumber *string `json:"issuance_guarantor_bank_number,omitempty"`

	// 出票保证人开户行名称。
	IssuanceGuarantorBankName *string `json:"issuance_guarantor_bank_name,omitempty"`

	// 大写票据金额。
	AmountInWords *string `json:"amount_in_words,omitempty"`

	// 小写票据金额。
	AmountInFigures *string `json:"amount_in_figures,omitempty"`

	// 承兑人全称。
	AcceptorFullName *string `json:"acceptor_full_name,omitempty"`

	// 承兑人账号。
	AcceptorAccount *string `json:"acceptor_account,omitempty"`

	// 承兑人开户行行号。
	AcceptorBankNumber *string `json:"acceptor_bank_number,omitempty"`

	// 承兑人开户行名称。
	AcceptorBankName *string `json:"acceptor_bank_name,omitempty"`

	// 交易合同号。
	ContractNumber *string `json:"contract_number,omitempty"`

	// 能否转让。
	Assignability *string `json:"assignability,omitempty"`

	// 出票人承诺。
	IssuerCommitment *string `json:"issuer_commitment,omitempty"`

	// 承兑人承诺。
	AcceptorCommitment *string `json:"acceptor_commitment,omitempty"`

	// 承兑日期。
	AcceptanceDate *string `json:"acceptance_date,omitempty"`

	// 承兑保证人名称。
	AcceptanceGuarantorName *string `json:"acceptance_guarantor_name,omitempty"`

	// 承兑保证人地址。
	AcceptanceGuarantorAddress *string `json:"acceptance_guarantor_address,omitempty"`

	// 承兑保证人账号。
	AcceptanceGuarantorAccount *string `json:"acceptance_guarantor_account,omitempty"`

	// 承兑保证日期。
	AcceptanceGuaranteeDate *string `json:"acceptance_guarantee_date,omitempty"`

	// 承兑保证人开户行行号。
	AcceptanceGuarantorBankNumber *string `json:"acceptance_guarantor_bank_number,omitempty"`

	// 承兑保证人开户行名称。
	AcceptanceGuarantorBankName *string `json:"acceptance_guarantor_bank_name,omitempty"`

	// 出票人评级主体。
	IssuerRatingEntity *string `json:"issuer_rating_entity,omitempty"`

	// 出票人信用等级。
	IssuerCreditRating *string `json:"issuer_credit_rating,omitempty"`

	// 出票人评级到期日。
	IssuerRatingDueDate *string `json:"issuer_rating_due_date,omitempty"`

	// 承兑人评级主体。
	AcceptorRatingEntity *string `json:"acceptor_rating_entity,omitempty"`

	// 承兑人信用等级。
	AcceptorCreditRating *string `json:"acceptor_credit_rating,omitempty"`

	// 承兑人评级到期日。
	AcceptorRatingDueDate *string `json:"acceptor_rating_due_date,omitempty"`

	// 票据包号。
	BillPackageNumber *string `json:"bill_package_number,omitempty"`

	// 备注。
	Remarks *string `json:"remarks,omitempty"`

	// 各个字段的置信度。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o AcceptanceBillResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptanceBillResult struct{}"
	}

	return strings.Join([]string{"AcceptanceBillResult", string(data)}, " ")
}
