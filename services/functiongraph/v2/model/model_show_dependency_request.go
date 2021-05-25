package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDependencyRequest struct {
	// 依赖包的ID。

	DependId string `json:"depend_id"`
}

func (o ShowDependencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDependencyRequest struct{}"
	}

	return strings.Join([]string{"ShowDependencyRequest", string(data)}, " ")
}
