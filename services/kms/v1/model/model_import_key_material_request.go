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
type ImportKeyMaterialRequest struct {
	VersionId string                        `json:"version_id"`
	Body      *ImportKeyMaterialRequestBody `json:"body,omitempty"`
}

func (o ImportKeyMaterialRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ImportKeyMaterialRequest", string(data)}, " ")
}
