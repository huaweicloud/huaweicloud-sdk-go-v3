package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrainTicketResult struct {

	// 火车票左上角的车票ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 检票口信息。
	CheckPort *string `json:"check_port,omitempty"`

	// 车次。
	TrainNumber *string `json:"train_number,omitempty"`

	// 始发站。
	DepartureStation *string `json:"departure_station,omitempty"`

	// 终点站。
	DestinationStation *string `json:"destination_station,omitempty"`

	// 始发站拼音。
	DepartureStationEn *string `json:"departure_station_en,omitempty"`

	// 终点站拼音。
	DestinationStationEn *string `json:"destination_station_en,omitempty"`

	// 开车时间。
	DepartureTime *string `json:"departure_time,omitempty"`

	// 座位号。
	SeatNumber *string `json:"seat_number,omitempty"`

	// 票价。
	TicketPrice *string `json:"ticket_price,omitempty"`

	// 售票方式。
	SaleMethod *string `json:"sale_method,omitempty"`

	// 座位类别。
	SeatCategory *string `json:"seat_category,omitempty"`

	// 是否改签票, \"Yes\"表示改签票，\"No\"表示非改签票。
	TicketChanging *string `json:"ticket_changing,omitempty"`

	// 车票持有人的身份证号。
	IdNumber *string `json:"id_number,omitempty"`

	// 车票持有人姓名。
	Name *string `json:"name,omitempty"`

	// 车票最下方的序列号。
	LogId *string `json:"log_id,omitempty"`

	// 车票售票地点。
	SaleLocation *string `json:"sale_location,omitempty"`

	// 类型。包含以下几种类型： - paper：纸质火车票 - electronic：电子发票 - refund_old：旧版退票凭证 - refund_new：新版退票凭证
	InvoiceStyle *string `json:"invoice_style,omitempty"`

	// 开票时间
	IssueDate *string `json:"issue_date,omitempty"`

	// 打折标识
	DiscountMark *string `json:"discount_mark,omitempty"`

	// 电子客票号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 税金价格
	TaxAmount *string `json:"tax_amount,omitempty"`

	// 税率
	TaxRate *string `json:"tax_rate,omitempty"`

	// 是否是空调车厢
	AirConditioning *string `json:"air_conditioning,omitempty"`

	// 原发票号码
	OriginalInvoiceNumber *string `json:"original_invoice_number,omitempty"`

	// 统一社会信用号码
	UnifiedSocialCreditCode *string `json:"unified_social_credit_code,omitempty"`

	// 购买方名称
	BuyerName *string `json:"buyer_name,omitempty"`

	// 不含税价格
	TotalAmountExcludingTax *string `json:"total_amount_excluding_tax,omitempty"`

	// 发票号码
	InvoiceNumber *string `json:"invoice_number,omitempty"`

	// 是否有印章，True表示有印章，False表示不含印章，字段默认为False
	SealMark *bool `json:"seal_mark,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 地区
	Area *string `json:"area,omitempty"`

	// 收据编码
	ReceiptNumber *string `json:"receipt_number,omitempty"`

	// 小写票据金额
	AmountInFigures *string `json:"amount_in_figures,omitempty"`

	// 大写票据金额
	AmountInWords *string `json:"amount_in_words,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o TrainTicketResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainTicketResult struct{}"
	}

	return strings.Join([]string{"TrainTicketResult", string(data)}, " ")
}
