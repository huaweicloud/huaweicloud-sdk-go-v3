package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThailandIdcardResult struct {

	// 标示正面还是反面，取值为front或back。
	Side *string `json:"side,omitempty" xml:"side"`

	// 身份证号。
	IdNumber *string `json:"id_number,omitempty" xml:"id_number"`

	// 泰文名字。
	NameTh *string `json:"name_th,omitempty" xml:"name_th"`

	// 英文名字。
	FirstNameEn *string `json:"first_name_en,omitempty" xml:"first_name_en"`

	// 英文姓氏。
	LastNameEn *string `json:"last_name_en,omitempty" xml:"last_name_en"`

	// 泰文出生日期。
	DateOfBirthTh *string `json:"date_of_birth_th,omitempty" xml:"date_of_birth_th"`

	// 英文出生日期。
	DateOfBirthEn *string `json:"date_of_birth_en,omitempty" xml:"date_of_birth_en"`

	// 宗教。
	ReligionTh *string `json:"religion_th,omitempty" xml:"religion_th"`

	// 地址。
	AddressTh *string `json:"address_th,omitempty" xml:"address_th"`

	// 泰文签发日期。
	DateOfIssueTh *string `json:"date_of_issue_th,omitempty" xml:"date_of_issue_th"`

	// 英文签发日期。
	DateOfIssueEn *string `json:"date_of_issue_en,omitempty" xml:"date_of_issue_en"`

	// 泰文有效期。
	DateOfExpiryTh *string `json:"date_of_expiry_th,omitempty" xml:"date_of_expiry_th"`

	// 英文有效期。
	DateOfExpiryEn *string `json:"date_of_expiry_en,omitempty" xml:"date_of_expiry_en"`

	// 序列号。
	SerialNumber *string `json:"serial_number,omitempty" xml:"serial_number"`

	// 身份证反面卡号。
	CardNumber *string `json:"card_number,omitempty" xml:"card_number"`

	// 激光码。
	LaserNumber *string `json:"laser_number,omitempty" xml:"laser_number"`

	Confidence *ThailandIdcardConfidence `json:"confidence,omitempty" xml:"confidence"`

	// 头像的base64编码。 当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty" xml:"portrait_image"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty" xml:"portrait_location"`

	// 身份证的类型。取值如下所示： - normal：身份证原件 - copy：复印的身份证 当输入参数“return_idcard_type”为“true”时，才返回该参数。
	IdcardType *string `json:"idcard_type,omitempty" xml:"idcard_type"`
}

func (o ThailandIdcardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThailandIdcardResult struct{}"
	}

	return strings.Join([]string{"ThailandIdcardResult", string(data)}, " ")
}
