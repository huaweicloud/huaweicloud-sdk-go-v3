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
type BatchCreateKmsTagsResponse struct {
}

func (o BatchCreateKmsTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateKmsTagsResponse", string(data)}, " ")
}
