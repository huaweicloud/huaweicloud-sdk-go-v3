package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarIdcardResult struct {

	// 标示正面还是反面，取值为front或back。
	Side *string `json:"side,omitempty" xml:"side"`

	// 身份证类型。取值如下所示： - new_version：新版身份证 - old_version：旧版
	Class *string `json:"class,omitempty" xml:"class"`

	// 身份证号码。
	NrcId *string `json:"nrc_id,omitempty" xml:"nrc_id"`

	// 签发日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 父亲名字。
	FatherName *string `json:"father_name,omitempty" xml:"father_name"`

	// 出生日期。
	Birth *string `json:"birth,omitempty" xml:"birth"`

	// 族群或宗教。
	BloodlinesReligion *string `json:"bloodlines_religion,omitempty" xml:"bloodlines_religion"`

	// 身高。
	Height *string `json:"height,omitempty" xml:"height"`

	// 血型。
	BloodGroup *string `json:"blood_group,omitempty" xml:"blood_group"`

	// 身份证的卡号（背面）。
	CardId *string `json:"card_id,omitempty" xml:"card_id"`

	// 背面的身份证号码。
	NrcIdBack *string `json:"nrc_id_back,omitempty" xml:"nrc_id_back"`

	// 职业。
	Profession *string `json:"profession,omitempty" xml:"profession"`

	// 地址。
	Address *string `json:"address,omitempty" xml:"address"`

	Confidence *MyanmarIdcardConfidence `json:"confidence,omitempty" xml:"confidence"`

	// 头像的base64编码。 当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty" xml:"portrait_image"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty" xml:"portrait_location"`

	// 身份证的类型。取值如下所示： - normal：身份证原件 - copy：复印的身份证 当输入参数“return_idcard_type”为“true”时，才返回该参数。
	IdcardType *string `json:"idcard_type,omitempty" xml:"idcard_type"`
}

func (o MyanmarIdcardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarIdcardResult struct{}"
	}

	return strings.Join([]string{"MyanmarIdcardResult", string(data)}, " ")
}
