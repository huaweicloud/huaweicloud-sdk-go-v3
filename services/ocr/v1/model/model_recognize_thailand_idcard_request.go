package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeThailandIdcardRequest struct {
	Body *ThailandIdcardRequestBody `json:"body,omitempty"`
}

func (o RecognizeThailandIdcardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandIdcardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeThailandIdcardRequest", string(data)}, " ")
}
