package model

import (
	"encoding/json"

	"strings"
)

type ApplyConfigurationRequestBody struct {
	// 实例ID列表对象。

	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequestBody", string(data)}, " ")
}
