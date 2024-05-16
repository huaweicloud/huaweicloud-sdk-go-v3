package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeBusinessCardResponse Response Object
type RecognizeBusinessCardResponse struct {
	Result *BusinessCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardResponse", string(data)}, " ")
}
