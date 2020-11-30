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
type ImportKeyMaterialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportKeyMaterialResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ImportKeyMaterialResponse", string(data)}, " ")
}
