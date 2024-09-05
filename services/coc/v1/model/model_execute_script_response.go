package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptResponse Response Object
type ExecuteScriptResponse struct {

	// execute_uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptResponse", string(data)}, " ")
}
