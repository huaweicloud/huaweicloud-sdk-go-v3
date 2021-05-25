package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClusterTagsRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](mrs_02_9001.xml)。

	ClusterId string `json:"cluster_id"`
}

func (o ListClusterTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterTagsRequest", string(data)}, " ")
}
