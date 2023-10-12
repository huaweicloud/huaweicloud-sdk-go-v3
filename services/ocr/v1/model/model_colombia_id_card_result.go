package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColombiaIdCardResult struct {

	// 证件图片正反面信息。可选值包括： - front: 证件图片正面 - back:  证件图片反面
	Side *string `json:"side,omitempty"`

	// 卡证编号。当响应字段\"side\"为front时，返回此字段。
	Number *string `json:"number,omitempty"`

	// 名。当响应字段\"side\"为front时，返回此字段。
	Name *string `json:"name,omitempty"`

	// 姓。当响应字段\"side\"为front时，返回此字段。
	LastName *string `json:"last_name,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 出生地。
	BirthPlace *string `json:"birth_place,omitempty"`

	// 性别。
	Gender *string `json:"gender,omitempty"`

	// 血型。
	BloodType *string `json:"blood_type,omitempty"`

	// 签发日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 签发地区。
	IssueAuthority *string `json:"issue_authority,omitempty"`

	// 身高。
	Height *string `json:"height,omitempty"`

	// 公民编码一。
	CitizenCode1 *string `json:"citizen_code1,omitempty"`

	// 公民编码二。
	CitizenCode2 *string `json:"citizen_code2,omitempty"`

	// 公民编码三。
	CitizenCode3 *string `json:"citizen_code3,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。注：置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence map[string]float32 `json:"confidence,omitempty"`
}

func (o ColombiaIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColombiaIdCardResult struct{}"
	}

	return strings.Join([]string{"ColombiaIdCardResult", string(data)}, " ")
}
