package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListScriptResultsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	ScriptName string `json:"script_name"`

	InstanceId string `json:"instance_id"`
}

func (o ListScriptResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResultsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptResultsRequest", string(data)}, " ")
}
