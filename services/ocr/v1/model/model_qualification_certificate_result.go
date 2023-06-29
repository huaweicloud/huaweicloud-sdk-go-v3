package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QualificationCertificateResult
type QualificationCertificateResult struct {

	// 身份证号（非必有，依赖对应从业资格证板式）。
	IdNumber *string `json:"id_number,omitempty"`

	// 考核时间（非必有，依赖对应从业资格证板式）。
	AssessmentDate *string `json:"assessment_date,omitempty"`

	// 从业资格证号。
	CertificateNumber *string `json:"certificate_number,omitempty"`

	// 档案号（非必有，依赖对应从业资格证板式）。
	FileNumber *string `json:"file_number,omitempty"`

	// 福路通号（非必有，依赖对应从业资格证板式）。
	UnionCardNumber *string `json:"union_card_number,omitempty"`

	// 继续教育信息（非必有，依赖对应从业资格证板式）。
	ContinuingEducationInfo *string `json:"continuing_education_info,omitempty"`

	// 性别（非必有，依赖对应从业资格证板式）。
	Sex *string `json:"sex,omitempty"`

	// 联系电话（非必有，依赖对应从业资格证板式）。
	PhoneNumber *string `json:"phone_number,omitempty"`

	// 登记时间（非必有，依赖对应从业资格证板式）。
	RegistrationDate *string `json:"registration_date,omitempty"`

	// 单位（非必有，依赖对应从业资格证板式）。
	WorkUnit *string `json:"work_unit,omitempty"`

	// 诚信考核信息（非必有，依赖对应从业资格证板式）。
	IntegrityAssessmentInfo *string `json:"integrity_assessment_info,omitempty"`

	// 国籍（非必有，依赖对应从业资格证板式）。
	Nationality *string `json:"nationality,omitempty"`

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 住址。
	Address *string `json:"address,omitempty"`

	// 准驾车型（非必有，依赖对应从业资格证板式）。
	DrivingClass *string `json:"driving_class,omitempty"`

	// 发证机关（非必有，依赖对应从业资格证板式）。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 出生日期（非必有，依赖对应从业资格证板式）。
	BirthDate *string `json:"birth_date,omitempty"`

	// 从业资格列表。
	QualificationCategoryList *[]QualificationCategory `json:"qualification_category_list,omitempty"`

	Confidence *QualificationConfidence `json:"confidence,omitempty"`
}

func (o QualificationCertificateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualificationCertificateResult struct{}"
	}

	return strings.Join([]string{"QualificationCertificateResult", string(data)}, " ")
}
