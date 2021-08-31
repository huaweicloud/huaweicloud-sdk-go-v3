package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeFlightItineraryRequest struct {
	Body *FlightItineraryRequestBody `json:"body,omitempty"`
}

func (o RecognizeFlightItineraryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeFlightItineraryRequest struct{}"
	}

	return strings.Join([]string{"RecognizeFlightItineraryRequest", string(data)}, " ")
}
