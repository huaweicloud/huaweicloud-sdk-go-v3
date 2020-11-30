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
type ListKeyDetailResponse struct {
	KeyInfo        *KeyDetails `json:"key_info,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListKeyDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeyDetailResponse", string(data)}, " ")
}
