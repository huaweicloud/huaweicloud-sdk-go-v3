package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 作业名称

	JobName string `json:"job_name"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
