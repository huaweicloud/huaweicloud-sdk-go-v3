package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClustersRequest struct {
}

func (o ListClustersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
