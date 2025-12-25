package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRetrieveScriptRequest Request Object
type DeleteRetrieveScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 检索脚本 ID
	RetrieveScriptId string `json:"retrieve_script_id"`
}

func (o DeleteRetrieveScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRetrieveScriptRequest struct{}"
	}

	return strings.Join([]string{"DeleteRetrieveScriptRequest", string(data)}, " ")
}
