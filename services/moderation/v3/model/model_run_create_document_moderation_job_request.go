package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateDocumentModerationJobRequest Request Object
type RunCreateDocumentModerationJobRequest struct {
	Body *DocumentCreateRequest `json:"body,omitempty"`
}

func (o RunCreateDocumentModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateDocumentModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCreateDocumentModerationJobRequest", string(data)}, " ")
}
