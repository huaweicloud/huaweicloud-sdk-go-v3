package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeCambodianIdCardResponse Response Object
type RecognizeCambodianIdCardResponse struct {
	Result *CambodianIdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeCambodianIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeCambodianIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeCambodianIdCardResponse", string(data)}, " ")
}
