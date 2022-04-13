package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaxiInvoiceResult struct {
	// 归属地区。

	Location *string `json:"location,omitempty"`
	// 发票代码。

	Code *string `json:"code,omitempty"`
	// 发票号码。

	Number *string `json:"number,omitempty"`
	// 电话（包括电话、监督电话）。

	PhoneNumber *string `json:"phone_number,omitempty"`
	// 单位。

	Company *string `json:"company,omitempty"`
	// 车号。

	TaxiNumber *string `json:"taxi_number,omitempty"`
	// 证号。

	CertificateNumber *string `json:"certificate_number,omitempty"`
	// 识别编号。

	IdentificationNumber *string `json:"identification_number,omitempty"`
	// 开票日期。

	Date *string `json:"date,omitempty"`
	// 上车时间。

	BoardingTime *string `json:"boarding_time,omitempty"`
	// 下车时间。

	AlightingTime *string `json:"alighting_time,omitempty"`
	// 时间(起止时间、上下车时间)。

	Time *string `json:"time,omitempty"`
	// 单价。

	UnitPrice *string `json:"unit_price,omitempty"`
	// 总里程。

	Distance *string `json:"distance,omitempty"`
	// 等候时间。

	WaitingTime *string `json:"waiting_time,omitempty"`
	// 金额。

	Fare *string `json:"fare,omitempty"`
	// 燃油附加费。

	FuelOilSurcharge *string `json:"fuel_oil_surcharge,omitempty"`
	// 电调费（预约费）。

	CallServiceSurcharge *string `json:"call_service_surcharge,omitempty"`
	// 实收金额。

	Total *string `json:"total,omitempty"`
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o TaxiInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaxiInvoiceResult struct{}"
	}

	return strings.Join([]string{"TaxiInvoiceResult", string(data)}, " ")
}
