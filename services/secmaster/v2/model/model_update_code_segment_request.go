package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeSegmentRequest Request Object
type UpdateCodeSegmentRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 代码片段 ID
	CodeSegmentId string `json:"code_segment_id"`

	Body *UpdateCodeSegmentRequestBody `json:"body,omitempty"`
}

func (o UpdateCodeSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeSegmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateCodeSegmentRequest", string(data)}, " ")
}
