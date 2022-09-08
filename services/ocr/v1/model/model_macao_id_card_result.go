package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MacaoIdCardResult struct {

	// 证件图片正反面信息。可选值包括： - \"front\": 证件图片为正面 - \"back\": 证件图片为反面
	Side *string `json:"side,omitempty"`

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 英文姓名，姓名单词之间使用空格进行间隔。
	NameEn *string `json:"name_en,omitempty"`

	// 性别，返回“男”或“女”。
	Sex *string `json:"sex,omitempty"`

	// 本次发证日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 证件有效期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 出生日期。
	BirthDate *string `json:"birth_date,omitempty"`

	// 首次发证日期。
	InitialIssueDate *string `json:"initial_issue_date,omitempty"`

	// 身高。
	Height *string `json:"height,omitempty"`

	// 身份证号。
	Number *string `json:"number,omitempty"`

	// 身份证上的字母代码，表示出生地等信息。
	Symbols *string `json:"symbols,omitempty"`

	// 身份证背面第一行机器码。
	MachineCode1 *string `json:"machine_code1,omitempty"`

	// 身份证背面第二行机器码。
	MachineCode2 *string `json:"machine_code2,omitempty"`

	// 身份证背面第三行机器码。
	MachineCode3 *string `json:"machine_code3,omitempty"`

	// 身份证头像照片的Base64编码。 若入参“return_portrait_image”字段缺失时，则此字段不存在。
	PortraitImage *string `json:"portrait_image,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。注：置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o MacaoIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MacaoIdCardResult struct{}"
	}

	return strings.Join([]string{"MacaoIdCardResult", string(data)}, " ")
}
