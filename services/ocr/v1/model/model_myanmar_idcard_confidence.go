package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarIdcardConfidence struct {

	// 身份证号码置信度。
	NrcId *float32 `json:"nrc_id,omitempty" xml:"nrc_id"`

	// 签发日期置信度。
	IssueDate *float32 `json:"issue_date,omitempty" xml:"issue_date"`

	// 姓名置信度。
	Name *float32 `json:"name,omitempty" xml:"name"`

	// 出生日期置信度。
	Birth *float32 `json:"birth,omitempty" xml:"birth"`

	// 族群或宗教置信度。
	BloodlinesReligion *float32 `json:"bloodlines_religion,omitempty" xml:"bloodlines_religion"`

	// 身高置信度。
	Height *float32 `json:"height,omitempty" xml:"height"`

	// 血型置信度。
	BloodGroup *float32 `json:"blood_group,omitempty" xml:"blood_group"`

	// 身份证的卡号（背面）置信度。
	CardId *float32 `json:"card_id,omitempty" xml:"card_id"`

	// 背面的身份证号码。
	NrcIdBack *float32 `json:"nrc_id_back,omitempty" xml:"nrc_id_back"`

	// 职业置信度。
	Profession *float32 `json:"profession,omitempty" xml:"profession"`

	// 地址置信度。
	Address *float32 `json:"address,omitempty" xml:"address"`
}

func (o MyanmarIdcardConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarIdcardConfidence struct{}"
	}

	return strings.Join([]string{"MyanmarIdcardConfidence", string(data)}, " ")
}
