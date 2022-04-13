package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeGeneralTextRequest struct {
	Body *GeneralTextRequestBody `json:"body,omitempty"`
}

func (o RecognizeGeneralTextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTextRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTextRequest", string(data)}, " ")
}
