package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeBusinessCardRequest struct {
	Body *BusinessCardRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeBusinessCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardRequest", string(data)}, " ")
}
