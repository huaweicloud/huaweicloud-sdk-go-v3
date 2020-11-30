/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTagsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTagsRequest", string(data)}, " ")
}
