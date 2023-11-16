package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeruIdCardResult struct {

	// 身份证号。
	CuiNumber *string `json:"cui_number,omitempty"`

	// 第一姓氏。
	FirstSurname *string `json:"first_surname,omitempty"`

	// 第二姓氏。
	SecondSurname *string `json:"second_surname,omitempty"`

	// 名。
	GivenName *string `json:"given_name,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 婚姻状况。
	MaritalStatus *string `json:"marital_status,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty"`

	// 发行日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 失效日期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 出生地编码。
	BirthPlace *string `json:"birth_place,omitempty"`

	// 投票组。
	VotingGroup *string `json:"voting_group,omitempty"`

	// 器官捐赠意愿。
	OrganDonation *string `json:"organ_donation,omitempty"`

	// 注册日期。
	RegistrationDate *string `json:"registration_date,omitempty"`

	// 头像的base64编码。当输入参数“return_portrait_image”为“true”时，才返回该参数。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 头像在原图上的位置。 当输入参数“return_portrait_location”为“true”时，才返回该参数。以列表形式显示，包含头像区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PortraitLocation *[][]int32 `json:"portrait_location,omitempty"`

	// 地址。
	Address *string `json:"address,omitempty"`

	// 大区。
	Department *string `json:"department,omitempty"`

	// 省。
	Province *string `json:"province,omitempty"`

	// 区。
	District *string `json:"district,omitempty"`

	// 备注。
	Remarks *string `json:"remarks,omitempty"`

	// 机器码第一行。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 机器码第二行。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 机器码第三行。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 是否重新登记过。可选值如下所示： -  true: 已重新登记过 -  false: 未重新登记过
	Duplicate *bool `json:"duplicate,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。注：置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence map[string]float32 `json:"confidence,omitempty"`
}

func (o PeruIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeruIdCardResult struct{}"
	}

	return strings.Join([]string{"PeruIdCardResult", string(data)}, " ")
}
