package model

import (
	"encoding/json"

	"strings"
)

type UpdateInstanceConfigurationRequestBody struct {
	// 参数值对象，用户基于默认参数模板自定义的参数值。

	Values map[string]string `json:"values"`
}

func (o UpdateInstanceConfigurationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequestBody", string(data)}, " ")
}
