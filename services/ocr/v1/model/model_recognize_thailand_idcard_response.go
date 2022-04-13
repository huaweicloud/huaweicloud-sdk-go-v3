package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeThailandIdcardResponse struct {
	Result         *ThailandIdcardResult `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o RecognizeThailandIdcardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandIdcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeThailandIdcardResponse", string(data)}, " ")
}
