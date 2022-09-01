package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeGeneralTextRequest struct {
	Body *GeneralTextRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeGeneralTextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTextRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTextRequest", string(data)}, " ")
}
