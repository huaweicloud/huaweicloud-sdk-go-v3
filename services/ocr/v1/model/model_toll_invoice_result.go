package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TollInvoiceResult struct {

	// 发票代码。
	Code *string `json:"code,omitempty"`

	// 发票号码。
	Number *string `json:"number,omitempty"`

	// 入口。
	Entry *string `json:"entry,omitempty"`

	// 出口。
	Exit *string `json:"exit,omitempty"`

	// 收费金额。
	Amount *string `json:"amount,omitempty"`

	// 收费员。
	Cashier *string `json:"cashier,omitempty"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty"`

	// 日期。
	Date *string `json:"date,omitempty"`

	// 时间。
	Time *string `json:"time,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o TollInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TollInvoiceResult struct{}"
	}

	return strings.Join([]string{"TollInvoiceResult", string(data)}, " ")
}
