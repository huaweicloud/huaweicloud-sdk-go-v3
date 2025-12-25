package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnalysisScriptRequest Request Object
type CreateAnalysisScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAnalysisScriptRequestBody `json:"body,omitempty"`
}

func (o CreateAnalysisScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalysisScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateAnalysisScriptRequest", string(data)}, " ")
}
