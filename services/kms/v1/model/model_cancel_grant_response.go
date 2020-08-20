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
type CancelGrantResponse struct {
}

func (o CancelGrantResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelGrantResponse", string(data)}, " ")
}
