package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExitEntryPermitResult struct {

	// 姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty" xml:"name_en"`

	// 性别。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date"`

	// 证件号。
	Number *string `json:"number,omitempty" xml:"number"`

	// 签发机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 签发地点。
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place"`

	// 有效期限。
	ValidPeriod *string `json:"valid_period,omitempty" xml:"valid_period"`

	// 机器码。
	MachineCode *string `json:"machine_code,omitempty" xml:"machine_code"`

	// 头像的base64编码。当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty" xml:"portrait_image"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty" xml:"portrait_location"`

	// 证件类型。可选值包括： - \"往来港澳通行证 \" - \"往来台湾通行证\"
	Type *string `json:"type,omitempty" xml:"type"`

	// 证件图片正反面信息。可选值包括： - \"front\"：证件图片为正面 - \"back\"：证件图片为反面
	Side *string `json:"side,omitempty" xml:"side"`

	EndorsementInfoHk *ExitEntryPermitEndorsementInfo `json:"endorsement_info_hk,omitempty" xml:"endorsement_info_hk"`

	EndorsementInfoMo *ExitEntryPermitEndorsementInfo `json:"endorsement_info_mo,omitempty" xml:"endorsement_info_mo"`

	EndorsementInfoTw *ExitEntryPermitEndorsementInfo `json:"endorsement_info_tw,omitempty" xml:"endorsement_info_tw"`

	Confidence *ExitEntryPermitConfidence `json:"confidence,omitempty" xml:"confidence"`
}

func (o ExitEntryPermitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExitEntryPermitResult struct{}"
	}

	return strings.Join([]string{"ExitEntryPermitResult", string(data)}, " ")
}
