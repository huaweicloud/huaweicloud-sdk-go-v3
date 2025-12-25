package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchLogsRequest Request Object
type ListSearchLogsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ListLogsRequestBody `json:"body,omitempty"`
}

func (o ListSearchLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchLogsRequest struct{}"
	}

	return strings.Join([]string{"ListSearchLogsRequest", string(data)}, " ")
}
