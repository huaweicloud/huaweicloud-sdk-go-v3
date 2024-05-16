package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeTrainTicketResponse Response Object
type RecognizeTrainTicketResponse struct {
	Result *TrainTicketResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeTrainTicketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTrainTicketResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTrainTicketResponse", string(data)}, " ")
}
