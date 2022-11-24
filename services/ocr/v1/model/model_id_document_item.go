package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdDocumentItem struct {

	// 证件签发国家或地区代码，命名遵循ISO-3166 3位代码。当前支持国家列表见表1。
	CountryRegion *string `json:"country_region,omitempty"`

	// 证件类型，可选值如下： - PP: passport，国际护照。 - DL: driving license，驾驶证。 - ID: identification card，各国颁发的身份证类型证件，比如身份证、选民卡、社保卡等。
	IdType *string `json:"id_type,omitempty"`

	// 证件正面或反面,可选值： - front: 正面，一般是包含人像的那面 - back: 背面 对于只有一面的卡证，返回front
	Side *string `json:"side,omitempty"`

	// 名
	FirstName *string `json:"first_name,omitempty"`

	// 姓氏
	LastName *string `json:"last_name,omitempty"`

	// 性别，可选值： M:男性 F:女性 X:中性
	Sex *string `json:"sex,omitempty"`

	// 持有人国籍
	Nationality *string `json:"nationality,omitempty"`

	// 生日，格式为yyyy-mm-dd
	BirthDate *string `json:"birth_date,omitempty"`

	// 签发日期，yyyy-mm-dd
	IssueDate *string `json:"issue_date,omitempty"`

	// 有效日期，yyyy-mm-dd
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 证件号码
	DocumentNumber *string `json:"document_number,omitempty"`

	// 持有人通讯地址
	Address *string `json:"address,omitempty"`

	// 签发机关
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 可选返回，证件头像图像base64编码
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 字段置信度，为0~1之间的小数，值越大，表明识别结果越可靠
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o IdDocumentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdDocumentItem struct{}"
	}

	return strings.Join([]string{"IdDocumentItem", string(data)}, " ")
}
