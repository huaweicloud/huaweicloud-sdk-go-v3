package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCloudPersistentVolumeClaimsRequest struct {
	Name string `json:"name"`

	Namespace string `json:"namespace"`

	DeleteVolume *string `json:"deleteVolume,omitempty"`

	StorageType *string `json:"storageType,omitempty"`

	XClusterID *string `json:"X-Cluster-ID,omitempty"`
}

func (o DeleteCloudPersistentVolumeClaimsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCloudPersistentVolumeClaimsRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudPersistentVolumeClaimsRequest", string(data)}, " ")
}
