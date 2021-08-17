package model

import (
	"encoding/json"

	"strings"
)

type ServicesInfo struct {
	// 服务id

	Id *int32 `json:"id,omitempty"`
	// 服务名称

	Name *string `json:"name,omitempty"`
}

func (o ServicesInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServicesInfo struct{}"
	}

	return strings.Join([]string{"ServicesInfo", string(data)}, " ")
}
