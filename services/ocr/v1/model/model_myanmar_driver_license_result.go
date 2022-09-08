package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarDriverLicenseResult struct {

	// 缅文驾驶证号。
	CardNumber *string `json:"card_number,omitempty"`

	// 英文驾驶证号。
	CardNumberEn *string `json:"card_number_en,omitempty"`

	// 缅文名字。
	Name *string `json:"name,omitempty"`

	// 英文名字。
	NameEn *string `json:"name_en,omitempty"`

	// 缅文nrc号码。
	NrcId *string `json:"nrc_id,omitempty"`

	// 英文nrc号码。
	NrcIdEn *string `json:"nrc_id_en,omitempty"`

	// 缅文出生日期。
	Birth *string `json:"Birth,omitempty"`

	// 英文出生日期。
	BirthEn *string `json:"birth_en,omitempty"`

	// 缅文血型。
	BloodGroup *string `json:"blood_group,omitempty"`

	// 英文血型。
	BloodGroupEn *string `json:"blood_group_en,omitempty"`

	// 缅文有效期。
	ExpiriedDate *string `json:"expiried_date,omitempty"`

	// 英文有效期。
	ExpiriedDateEn *string `json:"expiried_date_en,omitempty"`

	Confidence *MyanmarDriverLicenseConfidence `json:"confidence,omitempty"`
}

func (o MyanmarDriverLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarDriverLicenseResult struct{}"
	}

	return strings.Join([]string{"MyanmarDriverLicenseResult", string(data)}, " ")
}
