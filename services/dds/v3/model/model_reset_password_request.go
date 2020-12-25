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
type ResetPasswordRequest struct {
	InstanceId string                    `json:"instance_id"`
	Body       *ResetPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
