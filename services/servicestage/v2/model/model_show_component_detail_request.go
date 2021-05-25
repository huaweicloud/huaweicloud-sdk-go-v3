package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowComponentDetailRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
}

func (o ShowComponentDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowComponentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowComponentDetailRequest", string(data)}, " ")
}
