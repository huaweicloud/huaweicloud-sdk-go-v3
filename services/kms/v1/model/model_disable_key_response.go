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
type DisableKeyResponse struct {
	KeyInfo *KeyStatusInfo `json:"key_info,omitempty"`
}

func (o DisableKeyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisableKeyResponse", string(data)}, " ")
}
