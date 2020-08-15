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
type CreateRandomRequest struct {
	VersionId string                `json:"version_id"`
	Body      *GenRandomRequestBody `json:"body,omitempty"`
}

func (o CreateRandomRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateRandomRequest", string(data)}, " ")
}
