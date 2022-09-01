package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListScriptResultsRequest struct {
	ScriptName string `json:"script_name" xml:"script_name"`

	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o ListScriptResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResultsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptResultsRequest", string(data)}, " ")
}
