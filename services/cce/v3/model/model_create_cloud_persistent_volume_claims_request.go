package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCloudPersistentVolumeClaimsRequest struct {
	Namespace  string                 `json:"namespace"`
	XClusterID *string                `json:"X-Cluster-ID,omitempty"`
	Body       *PersistentVolumeClaim `json:"body,omitempty"`
}

func (o CreateCloudPersistentVolumeClaimsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCloudPersistentVolumeClaimsRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudPersistentVolumeClaimsRequest", string(data)}, " ")
}
