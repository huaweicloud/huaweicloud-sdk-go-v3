package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizePassportResponse Response Object
type RecognizePassportResponse struct {
	Result *PassportResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizePassportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePassportResponse struct{}"
	}

	return strings.Join([]string{"RecognizePassportResponse", string(data)}, " ")
}
