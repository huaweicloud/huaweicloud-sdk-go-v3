package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExitEntryPermitResult struct {

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 证件号。
	Number *string `json:"number,omitempty"`

	// 签发机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 签发地点。
	IssuePlace *string `json:"issue_place,omitempty"`

	// 有效期限。
	ValidPeriod *string `json:"valid_period,omitempty"`

	// 机器码。
	MachineCode *string `json:"machine_code,omitempty"`

	// 头像的base64编码。当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 证件类型。可选值包括： - \"往来港澳通行证 \" - \"往来台湾通行证\"
	Type *string `json:"type,omitempty"`

	// 证件图片正反面信息。可选值包括： - \"front\"：证件图片为正面 - \"back\"：证件图片为反面
	Side *string `json:"side,omitempty"`

	EndorsementInfoHk *ExitEntryPermitEndorsementInfo `json:"endorsement_info_hk,omitempty"`

	EndorsementInfoMo *ExitEntryPermitEndorsementInfo `json:"endorsement_info_mo,omitempty"`

	EndorsementInfoTw *ExitEntryPermitEndorsementInfo `json:"endorsement_info_tw,omitempty"`

	Confidence *ExitEntryPermitConfidence `json:"confidence,omitempty"`
}

func (o ExitEntryPermitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExitEntryPermitResult struct{}"
	}

	return strings.Join([]string{"ExitEntryPermitResult", string(data)}, " ")
}
