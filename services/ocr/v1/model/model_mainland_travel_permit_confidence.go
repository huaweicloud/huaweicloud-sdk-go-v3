package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MainlandTravelPermitConfidence struct {

	// 中文姓名的置信度。
	Name *float32 `json:"name,omitempty"`

	// 英文姓名的置信度。
	NameEn *float32 `json:"name_en,omitempty"`

	// 出生日期的置信度。
	BirthDate *float32 `json:"birth_date,omitempty"`

	// 性别的置信度。
	Sex *float32 `json:"sex,omitempty"`

	// 有效期限的置信度。
	ValidPeriod *float32 `json:"valid_period,omitempty"`

	// 签发机关的置信度。
	IssuingAuthority *float32 `json:"issuing_authority,omitempty"`

	// 证件号的置信度。
	Number *float32 `json:"number,omitempty"`

	// 签发地点的置信度。
	IssuePlace *float32 `json:"issue_place,omitempty"`

	// 签发次数的置信度。
	IssueTimes *float32 `json:"issue_times,omitempty"`

	// 证件类别的置信度。
	Type *float32 `json:"type,omitempty"`

	// 证件图片正反面信息的置信度。
	Side *float32 `json:"side,omitempty"`

	// 回乡证背面的香港/澳门/台湾身份证姓名的置信度。
	IdName *float32 `json:"id_name,omitempty"`

	// 回乡证背面的香港/澳门/台湾身份证号码的置信度。
	IdNumber *float32 `json:"id_number,omitempty"`

	// 机读码第一行的置信度。
	MachineCode1 *float32 `json:"machine_code1,omitempty"`

	// 机读码第二行的置信度。
	MachineCode2 *float32 `json:"machine_code2,omitempty"`

	// 机读码第三行的置信度。
	MachineCode3 *float32 `json:"machine_code3,omitempty"`
}

func (o MainlandTravelPermitConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MainlandTravelPermitConfidence struct{}"
	}

	return strings.Join([]string{"MainlandTravelPermitConfidence", string(data)}, " ")
}
