package model

import (
	"encoding/json"

	"strings"
)

// 资源结构体。
type ResourceDto struct {
	// 资源id。例如，要查询的资源类型为device，那么对应的资源id就是device_id。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ResourceDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
