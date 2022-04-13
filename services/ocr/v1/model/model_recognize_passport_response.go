package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizePassportResponse struct {
	Result         *PassportResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizePassportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePassportResponse struct{}"
	}

	return strings.Join([]string{"RecognizePassportResponse", string(data)}, " ")
}
