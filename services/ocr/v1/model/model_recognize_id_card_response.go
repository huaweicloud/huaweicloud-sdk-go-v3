package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeIdCardResponse struct {
	Result         *IdCardResult `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o RecognizeIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeIdCardResponse", string(data)}, " ")
}
