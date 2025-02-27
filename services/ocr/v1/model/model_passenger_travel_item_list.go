package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PassengerTravelItemList
type PassengerTravelItemList struct {

	// 出行人。
	TravelerName *string `json:"traveler_name,omitempty"`

	// 有效身份证件号。
	IdNumber *string `json:"id_number,omitempty"`

	// 出行日期。
	TravelDate *string `json:"travel_date,omitempty"`

	// 出发地。
	DepartureLocation *string `json:"departure_location,omitempty"`

	// 到达地。
	ArrivalLocation *string `json:"arrival_location,omitempty"`

	// 等级。
	Class *string `json:"class,omitempty"`

	// 交通工具类型。
	TransportationType *string `json:"transportation_type,omitempty"`
}

func (o PassengerTravelItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PassengerTravelItemList struct{}"
	}

	return strings.Join([]string{"PassengerTravelItemList", string(data)}, " ")
}
