package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRetrieveScriptRequest Request Object
type ShowRetrieveScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 检索脚本ID
	RetrieveScriptId string `json:"retrieve_script_id"`
}

func (o ShowRetrieveScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrieveScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowRetrieveScriptRequest", string(data)}, " ")
}
