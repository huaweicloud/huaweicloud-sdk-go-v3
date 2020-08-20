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
type UpdateKeyAliasResponse struct {
	KeyInfo *KeyAliasInfo `json:"key_info,omitempty"`
}

func (o UpdateKeyAliasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeyAliasResponse", string(data)}, " ")
}
