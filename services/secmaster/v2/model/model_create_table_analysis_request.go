package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableAnalysisRequest Request Object
type CreateTableAnalysisRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *ListTableLogsRequestBody `json:"body,omitempty"`
}

func (o CreateTableAnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableAnalysisRequest struct{}"
	}

	return strings.Join([]string{"CreateTableAnalysisRequest", string(data)}, " ")
}
