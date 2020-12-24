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
type ShowAddonInstanceRequest struct {
	ContentType string `json:"Content-Type"`
	Id          string `json:"id"`
	ClusterId   string `json:"cluster_id"`
}

func (o ShowAddonInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAddonInstanceRequest", string(data)}, " ")
}
