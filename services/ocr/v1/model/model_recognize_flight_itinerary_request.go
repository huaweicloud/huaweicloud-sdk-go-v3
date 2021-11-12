package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeFlightItineraryRequest struct {
	Body *FlightItineraryRequestBody `json:"body,omitempty"`
}

func (o RecognizeFlightItineraryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFlightItineraryRequest struct{}"
	}

	return strings.Join([]string{"RecognizeFlightItineraryRequest", string(data)}, " ")
}
