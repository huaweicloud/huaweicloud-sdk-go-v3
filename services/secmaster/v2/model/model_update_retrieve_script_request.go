package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRetrieveScriptRequest Request Object
type UpdateRetrieveScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 检索脚本ID
	RetrieveScriptId string `json:"retrieve_script_id"`

	Body *UpdateRetrieveScriptRequestBody `json:"body,omitempty"`
}

func (o UpdateRetrieveScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRetrieveScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateRetrieveScriptRequest", string(data)}, " ")
}
