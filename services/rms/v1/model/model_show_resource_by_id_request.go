/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowResourceByIdRequest struct {
	Provider   string `json:"provider"`
	Type       string `json:"type"`
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceByIdRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowResourceByIdRequest", string(data)}, " ")
}
