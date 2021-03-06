package model

import (
	"encoding/json"

	"strings"
)

//
type NodeCreateRequest struct {
	// API类型，固定值“Node”，该值不可修改。

	Kind string `json:"kind"`
	// API版本，固定值“v3”，该值不可修改。

	ApiVersion string `json:"apiVersion"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec"`
}

func (o NodeCreateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeCreateRequest struct{}"
	}

	return strings.Join([]string{"NodeCreateRequest", string(data)}, " ")
}
