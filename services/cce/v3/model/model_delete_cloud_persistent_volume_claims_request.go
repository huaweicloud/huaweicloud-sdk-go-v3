/*
 * cce
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCloudPersistentVolumeClaimsRequest struct {
	Name        string  `json:"name"`
	Namespace   string  `json:"namespace"`
	ContentType string  `json:"Content-Type"`
	XClusterID  *string `json:"X-Cluster-ID,omitempty"`
}

func (o DeleteCloudPersistentVolumeClaimsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteCloudPersistentVolumeClaimsRequest", string(data)}, " ")
}
