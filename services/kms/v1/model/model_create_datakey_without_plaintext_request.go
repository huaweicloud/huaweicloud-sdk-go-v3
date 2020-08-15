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
type CreateDatakeyWithoutPlaintextRequest struct {
	VersionId string                    `json:"version_id"`
	Body      *CreateDatakeyRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyWithoutPlaintextRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDatakeyWithoutPlaintextRequest", string(data)}, " ")
}
