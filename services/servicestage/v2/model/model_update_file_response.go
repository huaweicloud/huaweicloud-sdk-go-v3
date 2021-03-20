package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateFileResponse struct {
	// 文件路径。

	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFileResponse struct{}"
	}

	return strings.Join([]string{"UpdateFileResponse", string(data)}, " ")
}
