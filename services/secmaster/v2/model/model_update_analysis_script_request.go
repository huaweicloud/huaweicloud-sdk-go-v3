package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAnalysisScriptRequest Request Object
type UpdateAnalysisScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 分析脚本 ID
	AnalysisScriptId string `json:"analysis_script_id"`

	Body *UpdateAnalysisScriptRequestBody `json:"body,omitempty"`
}

func (o UpdateAnalysisScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnalysisScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateAnalysisScriptRequest", string(data)}, " ")
}
