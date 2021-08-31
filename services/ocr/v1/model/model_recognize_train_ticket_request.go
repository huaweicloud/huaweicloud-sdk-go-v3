package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeTrainTicketRequest struct {
	Body *TrainTicketRequestBody `json:"body,omitempty"`
}

func (o RecognizeTrainTicketRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTrainTicketRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTrainTicketRequest", string(data)}, " ")
}
