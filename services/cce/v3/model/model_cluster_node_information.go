package model

import (
	"encoding/json"

	"strings"
)

//
type ClusterNodeInformation struct {
	Metadata *ClusterNodeInformationMetadata `json:"metadata"`
}

func (o ClusterNodeInformation) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ClusterNodeInformation struct{}"
	}

	return strings.Join([]string{"ClusterNodeInformation", string(data)}, " ")
}
