package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateClusterResponse struct {
	// 集群名称

	Name *string `json:"name,omitempty"`
	// 集群ID

	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
