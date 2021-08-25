package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunCheckPictureRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`

	Body *DeletePictureReq `json:"body,omitempty"`
}

func (o RunCheckPictureRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCheckPictureRequest struct{}"
	}

	return strings.Join([]string{"RunCheckPictureRequest", string(data)}, " ")
}
