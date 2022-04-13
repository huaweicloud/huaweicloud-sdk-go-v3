package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThailandIdcardConfidence struct {
	// 身份证号置信度。

	IdNumber *float32 `json:"id_number,omitempty"`
	// 泰文名字置信度。

	NameTh *float32 `json:"name_th,omitempty"`
	// 英文名字置信度。

	FirstNameEn *float32 `json:"first_name_en,omitempty"`
	// 英文姓氏置信度。

	LastNameEn *float32 `json:"last_name_en,omitempty"`
	// 泰文出生日期置信度。

	DateOfBirthTh *float32 `json:"date_of_birth_th,omitempty"`
	// 英文出生日期置信度。

	DateOfBirthEn *float32 `json:"date_of_birth_en,omitempty"`
	// 宗教置信度。

	ReligionTh *float32 `json:"religion_th,omitempty"`
	// 地址置信度。

	AddressTh *float32 `json:"address_th,omitempty"`
	// 泰文签发日期置信度。

	DateOfIssueTh *float32 `json:"date_of_issue_th,omitempty"`
	// 英文签发日期置信度。

	DateOfIssueEn *float32 `json:"date_of_issue_en,omitempty"`
	// 泰文有效期置信度。

	DateOfExpiryTh *float32 `json:"date_of_expiry_th,omitempty"`
	// 英文有效期置信度。

	DateOfExpiryEn *float32 `json:"date_of_expiry_en,omitempty"`
	// 序列号置信度。

	SerialNumber *float32 `json:"serial_number,omitempty"`
	// 身份证反面卡号置信度。

	CardNumber *float32 `json:"card_number,omitempty"`
	// 激光码置信度。

	LaserNumber *float32 `json:"laser_number,omitempty"`
}

func (o ThailandIdcardConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThailandIdcardConfidence struct{}"
	}

	return strings.Join([]string{"ThailandIdcardConfidence", string(data)}, " ")
}
