package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableRequest Request Object
type UpdateTableRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *UpdateTableRequestBody `json:"body,omitempty"`
}

func (o UpdateTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateTableRequest", string(data)}, " ")
}
