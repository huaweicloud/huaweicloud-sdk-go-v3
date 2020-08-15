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
type CreateKeyRequest struct {
	VersionId string                `json:"version_id"`
	Body      *CreateKeyRequestBody `json:"body,omitempty"`
}

func (o CreateKeyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKeyRequest", string(data)}, " ")
}
