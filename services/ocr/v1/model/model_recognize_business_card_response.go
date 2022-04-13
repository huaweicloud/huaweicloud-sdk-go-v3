package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeBusinessCardResponse struct {
	Result         *BusinessCardResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardResponse", string(data)}, " ")
}
