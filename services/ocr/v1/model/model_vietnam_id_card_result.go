package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VietnamIdCardResult
type VietnamIdCardResult struct {

	// 返回证件正反面。字段值为“front”或“back”
	Side *string `json:"side,omitempty"`

	// 卡证编号。
	Number *string `json:"number,omitempty"`

	// 姓名。
	FullName *string `json:"full_name,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty"`

	// 籍贯。
	OriginPlace *string `json:"origin_place,omitempty"`

	// 居住地。
	ResidencePlace *string `json:"residence_place,omitempty"`

	// 有效日期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 个人识别。当参数“side”为back时，返回此字段。
	PersonalIdentification *string `json:"personal_identification,omitempty"`

	// 签发日期。当参数“side”为back时，返回此字段。
	IssueDate *string `json:"issue_date,omitempty"`

	// 身份证背面第一行机器码。当参数“side”为back时，返回此字段。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 身份证背面第二行机器码。当参数“side”为back时，返回此字段。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 身份证背面第三行机器码。当参数“side”为back时，返回此字段。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。注：置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`

	// 当参数return_portrait_image为true时，返回头像的base64编码。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 当参数return_portrait_location为true时，返回头像在原图上的位置，以列表形式表示，包含头像区域四个顶点的二维坐标（x,y）；坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 输入参数return_idcard_type为true时，返回身份证的类型：normal是身份证原件，copy是复印的身份证，screen是屏幕翻拍。
	IdcardType *string `json:"idcard_type,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o VietnamIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VietnamIdCardResult struct{}"
	}

	return strings.Join([]string{"VietnamIdCardResult", string(data)}, " ")
}
