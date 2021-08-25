package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunSearchPictureRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`

	Body *SearchPictureReq `json:"body,omitempty"`
}

func (o RunSearchPictureRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunSearchPictureRequest struct{}"
	}

	return strings.Join([]string{"RunSearchPictureRequest", string(data)}, " ")
}
