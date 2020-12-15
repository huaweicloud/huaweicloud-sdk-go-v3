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
type UpdateNodeResponse struct {
	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值“Node”，该值不可修改。
	Kind           *string       `json:"kind,omitempty"`
	Metadata       *NodeMetadata `json:"metadata,omitempty"`
	Spec           *V3NodeSpec   `json:"spec,omitempty"`
	Status         *V3NodeStatus `json:"status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateNodeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNodeResponse", string(data)}, " ")
}
