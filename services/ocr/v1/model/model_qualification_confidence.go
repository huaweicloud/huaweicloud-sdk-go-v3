package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualificationConfidence struct {

	// 身份证号置信度。
	IdNumber *float32 `json:"id_number,omitempty" xml:"id_number"`

	// 考核时间置信度。
	AssessmentDate *float32 `json:"assessment_date,omitempty" xml:"assessment_date"`

	// 从业资格证号置信度。
	CertificateNumber *float32 `json:"certificate_number,omitempty" xml:"certificate_number"`

	// 档案号置信度。
	FileNumber *float32 `json:"file_number,omitempty" xml:"file_number"`

	// 福路通号置信度。
	UnionCardNumber *float32 `json:"union_card_number,omitempty" xml:"union_card_number"`

	// 继续教育信息置信度。
	ContinuingEducationInfo *float32 `json:"continuing_education_info,omitempty" xml:"continuing_education_info"`

	// 性别置信度。
	Sex *float32 `json:"sex,omitempty" xml:"sex"`

	// 联系电话置信度。
	PhoneNumber *float32 `json:"phone_number,omitempty" xml:"phone_number"`

	// 登记时间置信度。
	RegistrationDate *float32 `json:"registration_date,omitempty" xml:"registration_date"`

	// 单位置信度。
	WorkUnit *float32 `json:"work_unit,omitempty" xml:"work_unit"`

	// 诚信考核信息置信度。
	IntegrityAssessmentInfo *float32 `json:"integrity_assessment_info,omitempty" xml:"integrity_assessment_info"`

	// 国籍置信度。
	Nationality *float32 `json:"nationality,omitempty" xml:"nationality"`

	// 姓名置信度。
	Name *float32 `json:"name,omitempty" xml:"name"`

	// 住址置信度。
	Address *float32 `json:"address,omitempty" xml:"address"`

	// 准驾车型置信度。
	DrivingClass *float32 `json:"driving_class,omitempty" xml:"driving_class"`

	// 发证机关置信度。
	IssuingAuthority *float32 `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 出生日期置信度。
	BirthDate *float32 `json:"birth_date,omitempty" xml:"birth_date"`

	// 从业资格列表置信度。
	QualificationCategoryList *[]QualificationCategoryConfidence `json:"qualification_category_list,omitempty" xml:"qualification_category_list"`
}

func (o QualificationConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualificationConfidence struct{}"
	}

	return strings.Join([]string{"QualificationConfidence", string(data)}, " ")
}
