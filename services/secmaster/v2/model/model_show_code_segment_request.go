package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCodeSegmentRequest Request Object
type ShowCodeSegmentRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 代码片段 ID
	CodeSegmentId string `json:"code_segment_id"`
}

func (o ShowCodeSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCodeSegmentRequest struct{}"
	}

	return strings.Join([]string{"ShowCodeSegmentRequest", string(data)}, " ")
}
