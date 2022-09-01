package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 机票行程列表。
type ItineraryList struct {

	// 始发站。
	DepartureStation *string `json:"departure_station,omitempty" xml:"departure_station"`

	// 目的站。
	DestinationStation *string `json:"destination_station,omitempty" xml:"destination_station"`

	// 承运人。
	Carrier *string `json:"carrier,omitempty" xml:"carrier"`

	// 航班号。
	Flight *string `json:"flight,omitempty" xml:"flight"`

	// 座位等级。
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class"`

	// 日期。
	Date *string `json:"date,omitempty" xml:"date"`

	// 时间。
	Time *string `json:"time,omitempty" xml:"time"`

	// 客票类别。
	FareBasis *string `json:"fare_basis,omitempty" xml:"fare_basis"`

	// 客票生效日期。
	EffectiveDate *string `json:"effective_date,omitempty" xml:"effective_date"`

	// 有效截止日期。
	ExpiryDate *string `json:"expiry_date,omitempty" xml:"expiry_date"`

	// 免费行李。
	BaggageAllowance *string `json:"baggage_allowance,omitempty" xml:"baggage_allowance"`
}

func (o ItineraryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItineraryList struct{}"
	}

	return strings.Join([]string{"ItineraryList", string(data)}, " ")
}
