/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type CreateKeyResponse struct {
	KeyInfo *KeKInfo `json:"key_info,omitempty"`
}

func (o CreateKeyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKeyResponse", string(data)}, " ")
}
