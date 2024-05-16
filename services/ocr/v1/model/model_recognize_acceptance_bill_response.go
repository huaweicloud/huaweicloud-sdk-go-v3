package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeAcceptanceBillResponse Response Object
type RecognizeAcceptanceBillResponse struct {
	Result *AcceptanceBillResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeAcceptanceBillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeAcceptanceBillResponse struct{}"
	}

	return strings.Join([]string{"RecognizeAcceptanceBillResponse", string(data)}, " ")
}
