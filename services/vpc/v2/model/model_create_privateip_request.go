/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type CreatePrivateipRequest struct {
	Body *CreatePrivateipRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePrivateipRequest", string(data)}, " ")
}
