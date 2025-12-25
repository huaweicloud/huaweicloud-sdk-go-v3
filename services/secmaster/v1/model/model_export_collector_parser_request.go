package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCollectorParserRequest Request Object
type ExportCollectorParserRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ExportParserDto `json:"body,omitempty"`
}

func (o ExportCollectorParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCollectorParserRequest struct{}"
	}

	return strings.Join([]string{"ExportCollectorParserRequest", string(data)}, " ")
}
