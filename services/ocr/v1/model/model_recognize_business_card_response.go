package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeBusinessCardResponse Response Object
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
