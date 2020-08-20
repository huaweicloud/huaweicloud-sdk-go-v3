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
type CancelSelfGrantResponse struct {
}

func (o CancelSelfGrantResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelSelfGrantResponse", string(data)}, " ")
}
