package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeVinRequest struct {
	Body *VinRequestBody `json:"body,omitempty"`
}

func (o RecognizeVinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVinRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVinRequest", string(data)}, " ")
}
