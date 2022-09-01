package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeTrainTicketRequest struct {
	Body *TrainTicketRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeTrainTicketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTrainTicketRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTrainTicketRequest", string(data)}, " ")
}
