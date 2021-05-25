package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateClusterTagRequest struct {
	// 集群ID。

	ClusterId string `json:"cluster_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreateClusterTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateClusterTagRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterTagRequest", string(data)}, " ")
}
