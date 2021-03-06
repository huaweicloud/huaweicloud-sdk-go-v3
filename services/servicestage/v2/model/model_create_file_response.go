package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateFileResponse struct {
	// 文件路径。

	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFileResponse struct{}"
	}

	return strings.Join([]string{"CreateFileResponse", string(data)}, " ")
}
