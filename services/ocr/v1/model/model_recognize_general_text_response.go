package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeGeneralTextResponse Response Object
type RecognizeGeneralTextResponse struct {
	Result *GeneralTextResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeGeneralTextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTextResponse struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTextResponse", string(data)}, " ")
}
