package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDetailsOfEnvironmentVariableV2Response struct {
	// 变量值

	VariableValue *string `json:"variable_value,omitempty"`
	// 环境编号

	EnvId *string `json:"env_id,omitempty"`
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
	// 环境变量编号

	Id *string `json:"id,omitempty"`
	// 变量名

	VariableName   *string `json:"variable_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailsOfEnvironmentVariableV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfEnvironmentVariableV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfEnvironmentVariableV2Response", string(data)}, " ")
}
