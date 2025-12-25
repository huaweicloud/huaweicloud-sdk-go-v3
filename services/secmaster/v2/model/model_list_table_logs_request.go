package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableLogsRequest Request Object
type ListTableLogsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *ListTableLogsRequestBody `json:"body,omitempty"`
}

func (o ListTableLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableLogsRequest struct{}"
	}

	return strings.Join([]string{"ListTableLogsRequest", string(data)}, " ")
}
