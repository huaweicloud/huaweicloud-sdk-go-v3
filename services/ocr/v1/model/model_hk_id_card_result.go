package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HkIdCardResult struct {

	// 中文姓名。
	Name *string `json:"name,omitempty"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty"`

	// 性别。  男： value值为：M 女： value值为：F
	Sex *string `json:"sex,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 身份证号。
	Number *string `json:"number,omitempty"`

	// 证件符号。
	Symbols *string `json:"symbols,omitempty"`

	// 中文姓名对应电码。
	NameTelegraphCode *string `json:"name_telegraph_code,omitempty"`

	// 是否永久性居民身份证。  永久：value值为true 非永久：value值为false
	Permanent *bool `json:"permanent,omitempty"`

	// 首次领用日期。
	InitialIssueDate *string `json:"initial_issue_date,omitempty"`

	// 签发日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 头像的base64编码，默认返回尺寸较大的头像。 当输入参数“return_portrait_image”为true时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 各个字段的置信度。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o HkIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HkIdCardResult struct{}"
	}

	return strings.Join([]string{"HkIdCardResult", string(data)}, " ")
}
