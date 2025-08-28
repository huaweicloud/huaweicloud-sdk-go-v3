package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentSegmentInfoResponse Response Object
type UpdateDocumentSegmentInfoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDocumentSegmentInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentSegmentInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDocumentSegmentInfoResponse", string(data)}, " ")
}
