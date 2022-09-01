package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeHandwritingRequest struct {
	Body *HandwritingRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeHandwritingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHandwritingRequest struct{}"
	}

	return strings.Join([]string{"RecognizeHandwritingRequest", string(data)}, " ")
}
