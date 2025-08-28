package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentSegmentParamResponse Response Object
type UpdateDocumentSegmentParamResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDocumentSegmentParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentSegmentParamResponse struct{}"
	}

	return strings.Join([]string{"UpdateDocumentSegmentParamResponse", string(data)}, " ")
}
