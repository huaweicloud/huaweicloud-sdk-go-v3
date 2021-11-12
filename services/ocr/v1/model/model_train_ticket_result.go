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
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o TrainTicketResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainTicketResult struct{}"
	}

	return strings.Join([]string{"TrainTicketResult", string(data)}, " ")
}
