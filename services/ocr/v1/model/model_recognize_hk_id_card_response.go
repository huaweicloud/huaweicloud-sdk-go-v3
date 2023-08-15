package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeHkIdCardResponse Response Object
type RecognizeHkIdCardResponse struct {
	Result         *HkIdCardResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeHkIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHkIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeHkIdCardResponse", string(data)}, " ")
}
