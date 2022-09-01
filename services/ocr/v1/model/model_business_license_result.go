package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BusinessLicenseResult struct {

	//   - 老版本营业执照对应注册号。  - 新三证合一版本营业执照对应社会保障号。
	RegistrationNumber *string `json:"registration_number,omitempty" xml:"registration_number"`

	// 企业名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 公司/企业类型/主体类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 住所/营业场所/企业住所。
	Address *string `json:"address,omitempty" xml:"address"`

	// 法定代表人/负责人。
	LegalRepresentative *string `json:"legal_representative,omitempty" xml:"legal_representative"`

	// 注册资本。
	RegisteredCapital *string `json:"registered_capital,omitempty" xml:"registered_capital"`

	// 成立日期。
	FoundDate *string `json:"found_date,omitempty" xml:"found_date"`

	// 营业期限。
	BusinessTerm *string `json:"business_term,omitempty" xml:"business_term"`

	// 经营范围。
	BusinessScope *string `json:"business_scope,omitempty" xml:"business_scope"`

	// 发照日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty" xml:"confidence"`
}

func (o BusinessLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessLicenseResult struct{}"
	}

	return strings.Join([]string{"BusinessLicenseResult", string(data)}, " ")
}
