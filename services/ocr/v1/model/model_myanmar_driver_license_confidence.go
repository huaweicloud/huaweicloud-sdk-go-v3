package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarDriverLicenseConfidence struct {
	// 缅文驾驶证号置信度。

	CardNumber *float32 `json:"card_number,omitempty"`
	// 英文驾驶证号置信度。

	CardNumberEn *float32 `json:"card_number_en,omitempty"`
	// 缅文名字置信度。

	Name *float32 `json:"name,omitempty"`
	// 英文名字置信度。

	NameEn *float32 `json:"name_en,omitempty"`
	// 缅文nrc号码置信度。

	NrcId *float32 `json:"nrc_id,omitempty"`
	// 英文nrc号码置信度。

	NrcIdEn *float32 `json:"nrc_id_en,omitempty"`
	// 缅文出生日期置信度。

	Birth *float32 `json:"Birth,omitempty"`
	// 英文出生日期置信度。

	BirthEn *float32 `json:"birth_en,omitempty"`
	// 缅文血型置信度。

	BloodGroup *float32 `json:"blood_group,omitempty"`
	// 英文血型置信度。

	BloodGroupEn *float32 `json:"blood_group_en,omitempty"`
	// 缅文有效期置信度。

	ExpiriedDate *float32 `json:"expiried_date,omitempty"`
	// 英文有效期置信度。

	ExpiriedDateEn *float32 `json:"expiried_date_en,omitempty"`
}

func (o MyanmarDriverLicenseConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarDriverLicenseConfidence struct{}"
	}

	return strings.Join([]string{"MyanmarDriverLicenseConfidence", string(data)}, " ")
}
