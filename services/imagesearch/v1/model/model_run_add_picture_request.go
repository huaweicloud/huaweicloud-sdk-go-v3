package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunAddPictureRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`

	Body *AddPictureRequestReq `json:"body,omitempty"`
}

func (o RunAddPictureRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunAddPictureRequest struct{}"
	}

	return strings.Join([]string{"RunAddPictureRequest", string(data)}, " ")
}
