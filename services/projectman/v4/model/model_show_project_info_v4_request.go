package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProjectInfoV4Request struct {
	// 项目ID

	ProjectId string `json:"project_id"`
}

func (o ShowProjectInfoV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectInfoV4Request struct{}"
	}

	return strings.Join([]string{"ShowProjectInfoV4Request", string(data)}, " ")
}
