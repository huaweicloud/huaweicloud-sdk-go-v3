package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessLicenseResult
type BusinessLicenseResult struct {

	// - 老版本营业执照对应注册号。 - 新三证合一版本营业执照对应社会保障号。
	RegistrationNumber *string `json:"registration_number,omitempty"`

	// 企业名称。
	Name *string `json:"name,omitempty"`

	// 公司/企业类型/主体类型/类型。
	Type *string `json:"type,omitempty"`

	// 住所/营业场所/企业住所/主要经营场所/经营场所。
	Address *string `json:"address,omitempty"`

	// 法定代表人/负责人/执行事务合伙人/投资人/经营者。
	LegalRepresentative *string `json:"legal_representative,omitempty"`

	// 注册资本/出资额。
	RegisteredCapital *string `json:"registered_capital,omitempty"`

	// 组成形式。
	OrganizationForm *string `json:"organization_form,omitempty"`

	// 成立日期/注册日期。
	FoundDate *string `json:"found_date,omitempty"`

	// 营业期限。
	BusinessTerm *string `json:"business_term,omitempty"`

	// 经营范围。
	BusinessScope *string `json:"business_scope,omitempty"`

	// 发照日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o BusinessLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessLicenseResult struct{}"
	}

	return strings.Join([]string{"BusinessLicenseResult", string(data)}, " ")
}
