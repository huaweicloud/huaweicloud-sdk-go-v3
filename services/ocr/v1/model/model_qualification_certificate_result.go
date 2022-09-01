package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type QualificationCertificateResult struct {

	// 身份证号（非必有，依赖对应从业资格证板式）。
	IdNumber *string `json:"id_number,omitempty" xml:"id_number"`

	// 考核时间（非必有，依赖对应从业资格证板式）。
	AssessmentDate *string `json:"assessment_date,omitempty" xml:"assessment_date"`

	// 从业资格证号。
	CertificateNumber *string `json:"certificate_number,omitempty" xml:"certificate_number"`

	// 档案号（非必有，依赖对应从业资格证板式）。
	FileNumber *string `json:"file_number,omitempty" xml:"file_number"`

	// 福路通号（非必有，依赖对应从业资格证板式）。
	UnionCardNumber *string `json:"union_card_number,omitempty" xml:"union_card_number"`

	// 继续教育信息（非必有，依赖对应从业资格证板式）。
	ContinuingEducationInfo *string `json:"continuing_education_info,omitempty" xml:"continuing_education_info"`

	// 性别（非必有，依赖对应从业资格证板式）。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 联系电话（非必有，依赖对应从业资格证板式）。
	PhoneNumber *string `json:"phone_number,omitempty" xml:"phone_number"`

	// 登记时间（非必有，依赖对应从业资格证板式）。
	RegistrationDate *string `json:"registration_date,omitempty" xml:"registration_date"`

	// 单位（非必有，依赖对应从业资格证板式）。
	WorkUnit *string `json:"work_unit,omitempty" xml:"work_unit"`

	// 诚信考核信息（非必有，依赖对应从业资格证板式）。
	IntegrityAssessmentInfo *string `json:"integrity_assessment_info,omitempty" xml:"integrity_assessment_info"`

	// 国籍（非必有，依赖对应从业资格证板式）。
	Nationality *string `json:"nationality,omitempty" xml:"nationality"`

	// 姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 住址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 准驾车型（非必有，依赖对应从业资格证板式）。
	DrivingClass *string `json:"driving_class,omitempty" xml:"driving_class"`

	// 发证机关（非必有，依赖对应从业资格证板式）。
	IssuingAuthority *string `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 出生日期（非必有，依赖对应从业资格证板式）。
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date"`

	// 从业资格列表。
	QualificationCategoryList *[]QualificationCategory `json:"qualification_category_list,omitempty" xml:"qualification_category_list"`

	Confidence *QualificationConfidence `json:"confidence,omitempty" xml:"confidence"`
}

func (o QualificationCertificateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualificationCertificateResult struct{}"
	}

	return strings.Join([]string{"QualificationCertificateResult", string(data)}, " ")
}
