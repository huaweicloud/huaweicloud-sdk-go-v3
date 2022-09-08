package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThailandIdcardResult struct {

	// 标示正面还是反面，取值为front或back。
	Side *string `json:"side,omitempty"`

	// 身份证号。
	IdNumber *string `json:"id_number,omitempty"`

	// 泰文名字。
	NameTh *string `json:"name_th,omitempty"`

	// 英文名字。
	FirstNameEn *string `json:"first_name_en,omitempty"`

	// 英文姓氏。
	LastNameEn *string `json:"last_name_en,omitempty"`

	// 泰文出生日期。
	DateOfBirthTh *string `json:"date_of_birth_th,omitempty"`

	// 英文出生日期。
	DateOfBirthEn *string `json:"date_of_birth_en,omitempty"`

	// 宗教。
	ReligionTh *string `json:"religion_th,omitempty"`

	// 地址。
	AddressTh *string `json:"address_th,omitempty"`

	// 泰文签发日期。
	DateOfIssueTh *string `json:"date_of_issue_th,omitempty"`

	// 英文签发日期。
	DateOfIssueEn *string `json:"date_of_issue_en,omitempty"`

	// 泰文有效期。
	DateOfExpiryTh *string `json:"date_of_expiry_th,omitempty"`

	// 英文有效期。
	DateOfExpiryEn *string `json:"date_of_expiry_en,omitempty"`

	// 序列号。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 身份证反面卡号。
	CardNumber *string `json:"card_number,omitempty"`

	// 激光码。
	LaserNumber *string `json:"laser_number,omitempty"`

	Confidence *ThailandIdcardConfidence `json:"confidence,omitempty"`

	// 头像的base64编码。 当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 身份证的类型。取值如下所示： - normal：身份证原件 - copy：复印的身份证 当输入参数“return_idcard_type”为“true”时，才返回该参数。
	IdcardType *string `json:"idcard_type,omitempty"`
}

func (o ThailandIdcardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThailandIdcardResult struct{}"
	}

	return strings.Join([]string{"ThailandIdcardResult", string(data)}, " ")
}
