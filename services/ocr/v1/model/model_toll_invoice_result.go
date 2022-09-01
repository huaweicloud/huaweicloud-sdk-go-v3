package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TollInvoiceResult struct {

	// 发票代码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 发票号码。
	Number *string `json:"number,omitempty" xml:"number"`

	// 入口。
	Entry *string `json:"entry,omitempty" xml:"entry"`

	// 出口。
	Exit *string `json:"exit,omitempty" xml:"exit"`

	// 收费金额。
	Amount *string `json:"amount,omitempty" xml:"amount"`

	// 收费员。
	Cashier *string `json:"cashier,omitempty" xml:"cashier"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty" xml:"vehicle_type"`

	// 日期。
	Date *string `json:"date,omitempty" xml:"date"`

	// 时间。
	Time *string `json:"time,omitempty" xml:"time"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty" xml:"confidence"`
}

func (o TollInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TollInvoiceResult struct{}"
	}

	return strings.Join([]string{"TollInvoiceResult", string(data)}, " ")
}
