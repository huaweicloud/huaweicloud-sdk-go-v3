package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalysisScriptRequest Request Object
type ShowAnalysisScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 分析脚本 ID
	AnalysisScriptId string `json:"analysis_script_id"`
}

func (o ShowAnalysisScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowAnalysisScriptRequest", string(data)}, " ")
}
