/*
 * CCE
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
type ShowClusterRequest struct {
	ClusterId   string  `json:"cluster_id"`
	ContentType string  `json:"Content-Type"`
	ErrorStatus *string `json:"errorStatus,omitempty"`
	Detail      *string `json:"detail,omitempty"`
}

func (o ShowClusterRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowClusterRequest", string(data)}, " ")
}
