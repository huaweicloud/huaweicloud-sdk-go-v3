package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunDeletePictureRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`

	Body *DeletePictureReq `json:"body,omitempty"`
}

func (o RunDeletePictureRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunDeletePictureRequest struct{}"
	}

	return strings.Join([]string{"RunDeletePictureRequest", string(data)}, " ")
}
