package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeFlightItineraryResponse struct {
	Result         *FlightItineraryResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RecognizeFlightItineraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFlightItineraryResponse struct{}"
	}

	return strings.Join([]string{"RecognizeFlightItineraryResponse", string(data)}, " ")
}
