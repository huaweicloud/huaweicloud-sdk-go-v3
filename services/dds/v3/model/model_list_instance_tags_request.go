/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstanceTagsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstanceTagsRequest", string(data)}, " ")
}
