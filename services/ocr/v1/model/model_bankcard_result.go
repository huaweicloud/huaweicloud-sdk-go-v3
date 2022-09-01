package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BankcardResult struct {

	// 发卡行。
	BankName *string `json:"bank_name,omitempty" xml:"bank_name"`

	// 银行卡号。
	CardNumber *string `json:"card_number,omitempty" xml:"card_number"`

	// 有效期开始日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 有效期截止日期。
	ExpiryDate *string `json:"expiry_date,omitempty" xml:"expiry_date"`

	// 银行卡类别，如：储蓄卡，信用卡。
	Type *string `json:"type,omitempty" xml:"type"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty" xml:"confidence"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty" xml:"text_location"`
}

func (o BankcardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BankcardResult struct{}"
	}

	return strings.Join([]string{"BankcardResult", string(data)}, " ")
}
