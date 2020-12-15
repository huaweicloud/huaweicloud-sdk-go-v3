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
type ListAddonInstancesRequest struct {
	ContentType string `json:"Content-Type"`
	ClusterId   string `json:"cluster_id"`
}

func (o ListAddonInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAddonInstancesRequest", string(data)}, " ")
}
