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
type EnableKeyRotationResponse struct {
}

func (o EnableKeyRotationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableKeyRotationResponse", string(data)}, " ")
}
