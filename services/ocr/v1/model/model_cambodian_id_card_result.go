package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CambodianIdCardResult struct {

	// 身份证号码。
	IdNumber *string `json:"id_number,omitempty"`

	// 高棉语版姓名。
	NameKh *string `json:"name_kh,omitempty"`

	// 英文姓名。
	NameEn *string `json:"name_en,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 身高。
	Height *string `json:"height,omitempty"`

	// 出生地。
	BirthPlace *string `json:"birth_place,omitempty"`

	// 地址，以空格分隔。
	Address *string `json:"address,omitempty"`

	// 签发起始日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 签发结束日期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 图片中的个人特征。
	Description *string `json:"description,omitempty"`

	// 机器码第一行。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 机器码第二行。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 机器码第三行。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 头像的base64编码。 当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 身份证的类型。当输入参数“ idcard_type ”为“true”时，才返回该参数。取值如下所示： - normal：身份证原件 - copy：复印的身份证
	IdcardType *string `json:"idcard_type,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o CambodianIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CambodianIdCardResult struct{}"
	}

	return strings.Join([]string{"CambodianIdCardResult", string(data)}, " ")
}
