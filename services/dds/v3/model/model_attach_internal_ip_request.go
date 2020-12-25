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
type AttachInternalIpRequest struct {
	InstanceId string                       `json:"instance_id"`
	Body       *AttachInternalIpRequestBody `json:"body,omitempty"`
}

func (o AttachInternalIpRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AttachInternalIpRequest", string(data)}, " ")
}
