package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeThailandIdcardResponse Response Object
type RecognizeThailandIdcardResponse struct {
	Result *ThailandIdcardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeThailandIdcardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandIdcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeThailandIdcardResponse", string(data)}, " ")
}
