package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeTrainTicketResponse struct {
	Result         *TrainTicketResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeTrainTicketResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTrainTicketResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTrainTicketResponse", string(data)}, " ")
}
