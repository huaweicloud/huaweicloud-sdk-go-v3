package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListClustersResponse struct {
	// 集群列表，请参见clusters参数说明

	Clusters       *[]Clusters `json:"clusters,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListClustersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClustersResponse struct{}"
	}

	return strings.Join([]string{"ListClustersResponse", string(data)}, " ")
}
