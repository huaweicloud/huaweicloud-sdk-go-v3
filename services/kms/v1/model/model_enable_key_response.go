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
type EnableKeyResponse struct {
	KeyInfo *KeyStatusInfo `json:"key_info,omitempty"`
}

func (o EnableKeyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableKeyResponse", string(data)}, " ")
}
