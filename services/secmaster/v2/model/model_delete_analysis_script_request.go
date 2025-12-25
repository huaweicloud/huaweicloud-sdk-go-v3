package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnalysisScriptRequest Request Object
type DeleteAnalysisScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 分析脚本 ID
	AnalysisScriptId string `json:"analysis_script_id"`
}

func (o DeleteAnalysisScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnalysisScriptRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnalysisScriptRequest", string(data)}, " ")
}
