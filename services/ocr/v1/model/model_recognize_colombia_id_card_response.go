package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeColombiaIdCardResponse Response Object
type RecognizeColombiaIdCardResponse struct {
	Result *ColombiaIdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeColombiaIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeColombiaIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeColombiaIdCardResponse", string(data)}, " ")
}
