package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteScriptResponse struct {
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptResponse", string(data)}, " ")
}
