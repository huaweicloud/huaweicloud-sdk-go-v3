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
type UpdateKeyRotationIntervalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateKeyRotationIntervalResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeyRotationIntervalResponse", string(data)}, " ")
}
