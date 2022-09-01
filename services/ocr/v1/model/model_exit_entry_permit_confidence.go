package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExitEntryPermitConfidence struct {

	// 姓名的置信度。
	Name *float32 `json:"name,omitempty" xml:"name"`

	// 英文姓名的置信度。
	NameEn *float32 `json:"name_en,omitempty" xml:"name_en"`

	// 出生日期的置信度。
	BirthDate *float32 `json:"birth_date,omitempty" xml:"birth_date"`

	// 性别的置信度
	Sex *float32 `json:"sex,omitempty" xml:"sex"`

	// 证件号的置信度。
	Number *float32 `json:"number,omitempty" xml:"number"`

	// 有效期限的置信度。
	ValidPeriod *float32 `json:"valid_period,omitempty" xml:"valid_period"`

	// 签发机关的置信度。
	IssuingAuthority *float32 `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 签发地点的置信度。
	IssuePlace *float32 `json:"issue_place,omitempty" xml:"issue_place"`

	// 机器码的置信度。
	MachineCode *float32 `json:"machine_code,omitempty" xml:"machine_code"`

	// 证件类型的置信度。
	Type *float32 `json:"type,omitempty" xml:"type"`

	// 证件图片正反面信息的置信度。
	Side *float32 `json:"side,omitempty" xml:"side"`

	// 香港签注信息的置信度。
	EndorsementInfoHk *interface{} `json:"endorsement_info_hk,omitempty" xml:"endorsement_info_hk"`

	// 澳门签注信息的置信度。
	EndorsementInfoMo *interface{} `json:"endorsement_info_mo,omitempty" xml:"endorsement_info_mo"`

	// 台湾签注信息的置信度。
	EndorsementInfoTw *interface{} `json:"endorsement_info_tw,omitempty" xml:"endorsement_info_tw"`
}

func (o ExitEntryPermitConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExitEntryPermitConfidence struct{}"
	}

	return strings.Join([]string{"ExitEntryPermitConfidence", string(data)}, " ")
}
