package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MainlandTravelPermitResult struct {

	// 中文姓名。
	Name *string `json:"name,omitempty"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 有效期限。
	ValidPeriod *string `json:"valid_period,omitempty"`

	// 签发机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 证件号。
	Number *string `json:"number,omitempty"`

	// 签发地点。
	IssuePlace *string `json:"issue_place,omitempty"`

	// 签发次数。
	IssueTimes *string `json:"issue_times,omitempty"`

	// 回乡证背面的香港/澳门/台湾身份证姓名。
	IdName *string `json:"id_name,omitempty"`

	// 回乡证背面的香港/澳门/台湾身份证号码。
	IdNumber *string `json:"id_number,omitempty"`

	// 机读码第一行。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 机读码第二行。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 机读码第三行。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 证件类别。可选值包括： - “港澳居民来往内地通行证” - “台湾居民来往大陆通行证”
	Type *string `json:"type,omitempty"`

	// 证件图片正反面信息。可选值包括： - \"front\"：证件图片为正面 - \"back\"：证件图片为反面
	Side *string `json:"side,omitempty"`

	// 头像的base64编码。当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	Confidence *MainlandTravelPermitConfidence `json:"confidence,omitempty"`
}

func (o MainlandTravelPermitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MainlandTravelPermitResult struct{}"
	}

	return strings.Join([]string{"MainlandTravelPermitResult", string(data)}, " ")
}
