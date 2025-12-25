package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeSegmentRequest Request Object
type CreateCodeSegmentRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateCodeSegmentRequestBody `json:"body,omitempty"`
}

func (o CreateCodeSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeSegmentRequest struct{}"
	}

	return strings.Join([]string{"CreateCodeSegmentRequest", string(data)}, " ")
}
