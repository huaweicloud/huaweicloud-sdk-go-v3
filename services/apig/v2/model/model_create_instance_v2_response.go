package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateInstanceV2Response struct {
	// 实例ID

	InstanceId *string `json:"instance_id,omitempty"`
	// 创建实例任务信息

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceV2Response struct{}"
	}

	return strings.Join([]string{"CreateInstanceV2Response", string(data)}, " ")
}
