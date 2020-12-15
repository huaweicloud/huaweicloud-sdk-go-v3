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
type ShowNodePoolRequest struct {
	ClusterId   string  `json:"cluster_id"`
	NodepoolId  string  `json:"nodepool_id"`
	ContentType string  `json:"Content-Type"`
	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o ShowNodePoolRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowNodePoolRequest", string(data)}, " ")
}
