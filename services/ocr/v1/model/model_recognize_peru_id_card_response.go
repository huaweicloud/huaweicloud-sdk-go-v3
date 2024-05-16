package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizePeruIdCardResponse Response Object
type RecognizePeruIdCardResponse struct {
	Result *PeruIdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizePeruIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePeruIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizePeruIdCardResponse", string(data)}, " ")
}
