package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentInstallScriptResponse Response Object
type ListAgentInstallScriptResponse struct {

	// agent安装脚本
	InstallScriptList *[]AgentInstallScriptResponseInfo `json:"install_script_list,omitempty"`
	HttpStatusCode    int                               `json:"-"`
}

func (o ListAgentInstallScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInstallScriptResponse struct{}"
	}

	return strings.Join([]string{"ListAgentInstallScriptResponse", string(data)}, " ")
}
