package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeIdCardResponse Response Object
type RecognizeIdCardResponse struct {
	Result *IdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeIdCardResponse", string(data)}, " ")
}
