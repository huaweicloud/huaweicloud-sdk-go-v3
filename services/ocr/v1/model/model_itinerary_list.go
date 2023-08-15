package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ItineraryList 机票行程列表。
type ItineraryList struct {

	// 始发站。
	DepartureStation *string `json:"departure_station,omitempty"`

	// 目的站。
	DestinationStation *string `json:"destination_station,omitempty"`

	// 承运人。
	Carrier *string `json:"carrier,omitempty"`

	// 航班号。
	Flight *string `json:"flight,omitempty"`

	// 座位等级。
	CabinClass *string `json:"cabin_class,omitempty"`

	// 日期。
	Date *string `json:"date,omitempty"`

	// 时间。
	Time *string `json:"time,omitempty"`

	// 客票类别。
	FareBasis *string `json:"fare_basis,omitempty"`

	// 客票生效日期。
	EffectiveDate *string `json:"effective_date,omitempty"`

	// 有效截止日期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 免费行李。
	BaggageAllowance *string `json:"baggage_allowance,omitempty"`
}

func (o ItineraryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItineraryList struct{}"
	}

	return strings.Join([]string{"ItineraryList", string(data)}, " ")
}
