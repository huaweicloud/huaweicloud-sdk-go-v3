package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableSchemaRequest Request Object
type UpdateTableSchemaRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *TableSchemaDto `json:"body,omitempty"`
}

func (o UpdateTableSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableSchemaRequest struct{}"
	}

	return strings.Join([]string{"UpdateTableSchemaRequest", string(data)}, " ")
}
