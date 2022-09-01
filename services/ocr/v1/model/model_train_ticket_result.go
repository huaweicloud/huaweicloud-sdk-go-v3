package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrainTicketResult struct {

	// 火车票左上角的车票ID。
	TicketId *string `json:"ticket_id,omitempty" xml:"ticket_id"`

	// 检票口信息。
	CheckPort *string `json:"check_port,omitempty" xml:"check_port"`

	// 车次。
	TrainNumber *string `json:"train_number,omitempty" xml:"train_number"`

	// 始发站。
	DepartureStation *string `json:"departure_station,omitempty" xml:"departure_station"`

	// 终点站。
	DestinationStation *string `json:"destination_station,omitempty" xml:"destination_station"`

	// 始发站拼音。
	DepartureStationEn *string `json:"departure_station_en,omitempty" xml:"departure_station_en"`

	// 终点站拼音。
	DestinationStationEn *string `json:"destination_station_en,omitempty" xml:"destination_station_en"`

	// 开车时间。
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time"`

	// 座位号。
	SeatNumber *string `json:"seat_number,omitempty" xml:"seat_number"`

	// 票价。
	TicketPrice *string `json:"ticket_price,omitempty" xml:"ticket_price"`

	// 售票方式。
	SaleMethod *string `json:"sale_method,omitempty" xml:"sale_method"`

	// 座位类别。
	SeatCategory *string `json:"seat_category,omitempty" xml:"seat_category"`

	// 是否改签票, \"Yes\"表示改签票，\"No\"表示非改签票。
	TicketChanging *string `json:"ticket_changing,omitempty" xml:"ticket_changing"`

	// 车票持有人的身份证号。
	IdNumber *string `json:"id_number,omitempty" xml:"id_number"`

	// 车票持有人姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 车票最下方的序列号。
	LogId *string `json:"log_id,omitempty" xml:"log_id"`

	// 车票售票地点。
	SaleLocation *string `json:"sale_location,omitempty" xml:"sale_location"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty" xml:"confidence"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty" xml:"text_location"`
}

func (o TrainTicketResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainTicketResult struct{}"
	}

	return strings.Join([]string{"TrainTicketResult", string(data)}, " ")
}
