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
type ListKeysRequest struct {
	VersionId string               `json:"version_id"`
	Body      *ListKeysRequestBody `json:"body,omitempty"`
}

func (o ListKeysRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeysRequest", string(data)}, " ")
}
