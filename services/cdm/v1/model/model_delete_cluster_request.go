package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteClusterRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`

	Body *CdmDeleteClusterReq `json:"body,omitempty"`
}

func (o DeleteClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
