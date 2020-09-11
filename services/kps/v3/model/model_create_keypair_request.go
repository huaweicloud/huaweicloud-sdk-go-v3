/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateKeypairRequest struct {
	Body *CreateKeypairRequestBody `json:"body,omitempty"`
}

func (o CreateKeypairRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKeypairRequest", string(data)}, " ")
}
