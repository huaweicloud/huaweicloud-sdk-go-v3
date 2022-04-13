package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeChileIdCardResponse struct {
	Result         *ChileIdCardResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeChileIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeChileIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeChileIdCardResponse", string(data)}, " ")
}
