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
type CreateNodePoolRequest struct {
	ContentType string    `json:"Content-Type"`
	ClusterId   string    `json:"cluster_id"`
	Body        *NodePool `json:"body,omitempty"`
}

func (o CreateNodePoolRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNodePoolRequest", string(data)}, " ")
}
