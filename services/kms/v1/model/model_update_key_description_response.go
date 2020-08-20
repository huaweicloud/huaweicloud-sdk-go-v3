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
type UpdateKeyDescriptionResponse struct {
	KeyInfo *KeyDescriptionInfo `json:"key_info,omitempty"`
}

func (o UpdateKeyDescriptionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeyDescriptionResponse", string(data)}, " ")
}
