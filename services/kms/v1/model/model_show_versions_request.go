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

// Request Object
type ShowVersionsRequest struct {
}

func (o ShowVersionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVersionsRequest", string(data)}, " ")
}
