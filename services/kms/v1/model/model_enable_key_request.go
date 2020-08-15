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
type EnableKeyRequest struct {
	VersionId string                 `json:"version_id"`
	Body      *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o EnableKeyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableKeyRequest", string(data)}, " ")
}
