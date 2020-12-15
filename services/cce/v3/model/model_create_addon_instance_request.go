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
type CreateAddonInstanceRequest struct {
	ContentType string           `json:"Content-Type"`
	Body        *InstanceRequest `json:"body,omitempty"`
}

func (o CreateAddonInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAddonInstanceRequest", string(data)}, " ")
}
