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
type CreateKmsTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateKmsTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKmsTagResponse", string(data)}, " ")
}
