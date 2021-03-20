package model

import (
	"encoding/json"

	"strings"
)

type FilePath struct {
	// 文件在OBS上的路径

	Path *string `json:"path,omitempty"`
}

func (o FilePath) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FilePath struct{}"
	}

	return strings.Join([]string{"FilePath", string(data)}, " ")
}
