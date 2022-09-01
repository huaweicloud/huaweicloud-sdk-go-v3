package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteScriptResponse struct {
	InstanceId     *string `json:"instanceId,omitempty" xml:"instanceId"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptResponse", string(data)}, " ")
}
