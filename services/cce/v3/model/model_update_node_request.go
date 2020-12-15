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
type UpdateNodeRequest struct {
	ClusterId   string                     `json:"cluster_id"`
	NodeId      string                     `json:"node_id"`
	ContentType string                     `json:"Content-Type"`
	ErrorStatus *string                    `json:"errorStatus,omitempty"`
	Body        *CceClusterNodeInformation `json:"body,omitempty"`
}

func (o UpdateNodeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNodeRequest", string(data)}, " ")
}
