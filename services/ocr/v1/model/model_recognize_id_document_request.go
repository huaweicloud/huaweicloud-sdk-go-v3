package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeIdDocumentRequest struct {
	Body *IdDocumentRequestBody `json:"body,omitempty"`
}

func (o RecognizeIdDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdDocumentRequest struct{}"
	}

	return strings.Join([]string{"RecognizeIdDocumentRequest", string(data)}, " ")
}
