package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarDriverLicenseResult struct {

	// 缅文驾驶证号。
	CardNumber *string `json:"card_number,omitempty" xml:"card_number"`

	// 英文驾驶证号。
	CardNumberEn *string `json:"card_number_en,omitempty" xml:"card_number_en"`

	// 缅文名字。
	Name *string `json:"name,omitempty" xml:"name"`

	// 英文名字。
	NameEn *string `json:"name_en,omitempty" xml:"name_en"`

	// 缅文nrc号码。
	NrcId *string `json:"nrc_id,omitempty" xml:"nrc_id"`

	// 英文nrc号码。
	NrcIdEn *string `json:"nrc_id_en,omitempty" xml:"nrc_id_en"`

	// 缅文出生日期。
	Birth *string `json:"Birth,omitempty" xml:"Birth"`

	// 英文出生日期。
	BirthEn *string `json:"birth_en,omitempty" xml:"birth_en"`

	// 缅文血型。
	BloodGroup *string `json:"blood_group,omitempty" xml:"blood_group"`

	// 英文血型。
	BloodGroupEn *string `json:"blood_group_en,omitempty" xml:"blood_group_en"`

	// 缅文有效期。
	ExpiriedDate *string `json:"expiried_date,omitempty" xml:"expiried_date"`

	// 英文有效期。
	ExpiriedDateEn *string `json:"expiried_date_en,omitempty" xml:"expiried_date_en"`

	Confidence *MyanmarDriverLicenseConfidence `json:"confidence,omitempty" xml:"confidence"`
}

func (o MyanmarDriverLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarDriverLicenseResult struct{}"
	}

	return strings.Join([]string{"MyanmarDriverLicenseResult", string(data)}, " ")
}
