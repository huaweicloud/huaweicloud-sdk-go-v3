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

type V3NodeCreateRequest struct {
	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string `json:"apiVersion"`
	// API类型，固定值“Node”，该值不可修改。
	Kind     string        `json:"kind"`
	Metadata *NodeMetadata `json:"metadata,omitempty"`
	Spec     *V3NodeSpec   `json:"spec"`
}

func (o V3NodeCreateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"V3NodeCreateRequest", string(data)}, " ")
}
