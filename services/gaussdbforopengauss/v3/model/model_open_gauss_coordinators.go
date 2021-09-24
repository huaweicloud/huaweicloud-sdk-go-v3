package model

import (
	"encoding/json"

	"strings"
)

// CN横向扩容时必填
type OpenGaussCoordinators struct {
	// 新增CN横向扩容每个节点的可用区

	AzCode string `json:"az_code"`
}

func (o OpenGaussCoordinators) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussCoordinators struct{}"
	}

	return strings.Join([]string{"OpenGaussCoordinators", string(data)}, " ")
}
