package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThailandIdcardConfidence struct {

	// 身份证号置信度。
	IdNumber *float32 `json:"id_number,omitempty" xml:"id_number"`

	// 泰文名字置信度。
	NameTh *float32 `json:"name_th,omitempty" xml:"name_th"`

	// 英文名字置信度。
	FirstNameEn *float32 `json:"first_name_en,omitempty" xml:"first_name_en"`

	// 英文姓氏置信度。
	LastNameEn *float32 `json:"last_name_en,omitempty" xml:"last_name_en"`

	// 泰文出生日期置信度。
	DateOfBirthTh *float32 `json:"date_of_birth_th,omitempty" xml:"date_of_birth_th"`

	// 英文出生日期置信度。
	DateOfBirthEn *float32 `json:"date_of_birth_en,omitempty" xml:"date_of_birth_en"`

	// 宗教置信度。
	ReligionTh *float32 `json:"religion_th,omitempty" xml:"religion_th"`

	// 地址置信度。
	AddressTh *float32 `json:"address_th,omitempty" xml:"address_th"`

	// 泰文签发日期置信度。
	DateOfIssueTh *float32 `json:"date_of_issue_th,omitempty" xml:"date_of_issue_th"`

	// 英文签发日期置信度。
	DateOfIssueEn *float32 `json:"date_of_issue_en,omitempty" xml:"date_of_issue_en"`

	// 泰文有效期置信度。
	DateOfExpiryTh *float32 `json:"date_of_expiry_th,omitempty" xml:"date_of_expiry_th"`

	// 英文有效期置信度。
	DateOfExpiryEn *float32 `json:"date_of_expiry_en,omitempty" xml:"date_of_expiry_en"`

	// 序列号置信度。
	SerialNumber *float32 `json:"serial_number,omitempty" xml:"serial_number"`

	// 身份证反面卡号置信度。
	CardNumber *float32 `json:"card_number,omitempty" xml:"card_number"`

	// 激光码置信度。
	LaserNumber *float32 `json:"laser_number,omitempty" xml:"laser_number"`
}

func (o ThailandIdcardConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThailandIdcardConfidence struct{}"
	}

	return strings.Join([]string{"ThailandIdcardConfidence", string(data)}, " ")
}
