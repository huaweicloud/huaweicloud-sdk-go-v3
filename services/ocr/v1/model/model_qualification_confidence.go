package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualificationConfidence struct {

	// 身份证号置信度。
	IdNumber *float32 `json:"id_number,omitempty"`

	// 考核时间置信度。
	AssessmentDate *float32 `json:"assessment_date,omitempty"`

	// 从业资格证号置信度。
	CertificateNumber *float32 `json:"certificate_number,omitempty"`

	// 档案号置信度。
	FileNumber *float32 `json:"file_number,omitempty"`

	// 福路通号置信度。
	UnionCardNumber *float32 `json:"union_card_number,omitempty"`

	// 继续教育信息置信度。
	ContinuingEducationInfo *float32 `json:"continuing_education_info,omitempty"`

	// 性别置信度。
	Sex *float32 `json:"sex,omitempty"`

	// 联系电话置信度。
	PhoneNumber *float32 `json:"phone_number,omitempty"`

	// 登记时间置信度。
	RegistrationDate *float32 `json:"registration_date,omitempty"`

	// 单位置信度。
	WorkUnit *float32 `json:"work_unit,omitempty"`

	// 诚信考核信息置信度。
	IntegrityAssessmentInfo *float32 `json:"integrity_assessment_info,omitempty"`

	// 国籍置信度。
	Nationality *float32 `json:"nationality,omitempty"`

	// 姓名置信度。
	Name *float32 `json:"name,omitempty"`

	// 住址置信度。
	Address *float32 `json:"address,omitempty"`

	// 准驾车型置信度。
	DrivingClass *float32 `json:"driving_class,omitempty"`

	// 发证机关置信度。
	IssuingAuthority *float32 `json:"issuing_authority,omitempty"`

	// 出生日期置信度。
	BirthDate *float32 `json:"birth_date,omitempty"`

	// 从业资格列表置信度。
	QualificationCategoryList *[]QualificationCategoryConfidence `json:"qualification_category_list,omitempty"`
}

func (o QualificationConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualificationConfidence struct{}"
	}

	return strings.Join([]string{"QualificationConfidence", string(data)}, " ")
}
