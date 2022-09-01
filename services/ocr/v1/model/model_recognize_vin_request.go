package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeVinRequest struct {
	Body *VinRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeVinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVinRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVinRequest", string(data)}, " ")
}
