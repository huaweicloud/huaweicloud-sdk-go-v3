package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeFlightItineraryResponse struct {
	Result         *FlightItineraryResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RecognizeFlightItineraryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeFlightItineraryResponse struct{}"
	}

	return strings.Join([]string{"RecognizeFlightItineraryResponse", string(data)}, " ")
}
