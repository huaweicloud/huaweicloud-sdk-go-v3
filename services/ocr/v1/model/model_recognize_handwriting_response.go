package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeHandwritingResponse struct {
	Result         *HandwritingResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeHandwritingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHandwritingResponse struct{}"
	}

	return strings.Join([]string{"RecognizeHandwritingResponse", string(data)}, " ")
}
