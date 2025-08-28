package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewDocumentSegmentResponse Response Object
type PreviewDocumentSegmentResponse struct {
	Body *[]DocumentSegmentInfo `json:"body,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PreviewDocumentSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewDocumentSegmentResponse struct{}"
	}

	return strings.Join([]string{"PreviewDocumentSegmentResponse", string(data)}, " ")
}
