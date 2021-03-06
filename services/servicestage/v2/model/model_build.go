package model

import (
	"encoding/json"

	"strings"
)

// 构建工程。
type Build struct {
	Parameters *BuildInfoParameters `json:"parameters,omitempty"`
}

func (o Build) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Build struct{}"
	}

	return strings.Join([]string{"Build", string(data)}, " ")
}
