package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopScriptExecutionResponse Response Object
type StopScriptExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopScriptExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopScriptExecutionResponse struct{}"
	}

	return strings.Join([]string{"StopScriptExecutionResponse", string(data)}, " ")
}
