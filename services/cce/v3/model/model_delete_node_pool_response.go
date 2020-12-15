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

// Response Object
type DeleteNodePoolResponse struct {
	// API版本，固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值“NodePool”。
	Kind           *string           `json:"kind,omitempty"`
	Metadata       *NodePoolMetadata `json:"metadata,omitempty"`
	Spec           *NodePoolSpec     `json:"spec,omitempty"`
	Status         *NodePoolStatus   `json:"status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteNodePoolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteNodePoolResponse", string(data)}, " ")
}
